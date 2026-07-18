package socks

import (
        "encoding/binary"
        "io"
        "net"
        "sync"
        "time"
)

type udpSession struct {
        srcAddr   *net.UDPAddr // client address (phone via Streisand)
        dstAddr   *net.UDPAddr // destination address (peer)
        outConn   *net.UDPConn // single outbound connection
        lastUsed  time.Time
}

func (s *Server) handleUDPAssociate(client net.Conn) error {

        udp, err := net.ListenUDP(
                "udp",
                &net.UDPAddr{
                        IP:   net.ParseIP("127.0.0.1"),
                        Port: 0,
                },
        )
        if err != nil {
                return err
        }

        addr := udp.LocalAddr().(*net.UDPAddr)

        reply := []byte{
                0x05,
                0x00,
                0x00,
                0x01,
                127, 0, 0, 1,
                byte(addr.Port >> 8),
                byte(addr.Port),
        }

        if _, err := client.Write(reply); err != nil {
                udp.Close()
                return err
        }

        go s.relayUDP(udp)

        io.Copy(io.Discard, client)

        udp.Close()

        return nil
}

func (s *Server) relayUDP(conn *net.UDPConn) {

        buf := make([]byte, 65535)

        // Single outbound socket for all connections
        outSocket, err := net.ListenUDP("udp", nil)
        if err != nil {
                return
        }
        defer outSocket.Close()

        sessions := make(map[string]*udpSession)
        var mu sync.Mutex

        // Read responses from outSocket and forward to clients
        go func() {
                respBuf := make([]byte, 65535)
                for {
                        rn, remoteAddr, err := outSocket.ReadFromUDP(respBuf)
                        if err != nil {
                                return
                        }

                        mu.Lock()
                        session, ok := sessions[remoteAddr.String()]
                        if ok {
                                session.lastUsed = time.Now()
                        }
                        mu.Unlock()

                        if !ok {
                                continue
                        }

                        // Build SOCKS5 UDP response header
                        resp := make([]byte, 3+1+4+2+rn)
                        resp[0] = 0x00
                        resp[1] = 0x00
                        resp[2] = 0x00
                        resp[3] = 0x01
                        copy(resp[4:8], remoteAddr.IP.To4())
                        binary.BigEndian.PutUint16(resp[8:10], uint16(remoteAddr.Port))
                        copy(resp[10:], respBuf[:rn])

                        conn.WriteToUDP(resp, session.srcAddr)
                }
        }()

        // Cleanup old sessions
        go func() {
                for {
                        time.Sleep(30 * time.Second)
                        mu.Lock()
                        for key, session := range sessions {
                                if time.Since(session.lastUsed) > 120*time.Second {
                                        delete(sessions, key)
                                }
                        }
                        mu.Unlock()
                }
        }()

        for {
                n, src, err := conn.ReadFromUDP(buf)
                if err != nil {
                        return
                }

                // SOCKS5 UDP header: RSV(2) + FRAG(1) + ATYP(1) + DST.ADDR + DST.PORT + DATA
                if n < 10 {
                        continue
                }

                pos := 3
                atyp := buf[pos]
                pos++

                var dstIP net.IP
                var dstPort int

                switch atyp {
                case 0x01: // IPv4
                        if n < pos+4+2 {
                                continue
                        }
                        dstIP = net.IP(buf[pos : pos+4])
                        pos += 4
                case 0x03: // Domain
                        if n < pos+1 {
                                continue
                        }
                        domainLen := int(buf[pos])
                        pos++
                        if n < pos+domainLen+2 {
                                continue
                        }
                        domain := string(buf[pos : pos+domainLen])
                        pos += domainLen
                        ips, err := net.LookupIP(domain)
                        if err != nil || len(ips) == 0 {
                                continue
                        }
                        dstIP = ips[0]
                case 0x04: // IPv6
                        if n < pos+16+2 {
                                continue
                        }
                        dstIP = net.IP(buf[pos : pos+16])
                        pos += 16
                default:
                        continue
                }

                dstPort = int(binary.BigEndian.Uint16(buf[pos : pos+2]))
                pos += 2

                data := buf[pos:n]

                if dstIP == nil {
                        continue
                }

                dstAddr := &net.UDPAddr{
                        IP:   dstIP,
                        Port: dstPort,
                }

                key := dstAddr.String()

                mu.Lock()
                session, ok := sessions[key]
                if !ok {
                        session = &udpSession{
                                srcAddr:  src,
                                dstAddr:  dstAddr,
                                outConn:  outSocket,
                                lastUsed: time.Now(),
                        }
                        sessions[key] = session
                }
                session.lastUsed = time.Now()
                mu.Unlock()

                // Send to destination via single outSocket
                _, err = outSocket.WriteToUDP(data, dstAddr)
                if err != nil {
                        mu.Lock()
                        delete(sessions, key)
                        mu.Unlock()
                }
        }
}
