package socks

import (
	"log"
	"context"
	"encoding/binary"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"github.com/soulcpd1/xrm2/internal/outbound"
)

type udpSession struct {
	srcAddr   *net.UDPAddr
	dstAddr   *net.UDPAddr
	packetConn net.PacketConn
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
		0x05, 0x00, 0x00, 0x01,
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

	// Создаём агент для UDP через VPS (аналогично TCP)
	agent := outbound.NewAgent(os.Getenv("XRM_AGENT"))

	sessions := make(map[string]*udpSession)
	var mu sync.Mutex

	// Очистка старых сессий
	go func() {
		for {
			time.Sleep(30 * time.Second)
			mu.Lock()
			for key, session := range sessions {
				if time.Since(session.lastUsed) > 120*time.Second {
					if session.packetConn != nil {
						session.packetConn.Close()
					}
					delete(sessions, key)
				}
			}
			mu.Unlock()
		}
	}()

	for {
		n, src, err := conn.ReadFromUDP(buf)
		log.Printf("[UDP] received %d bytes from %s", n, src)
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

		dstAddr := &net.UDPAddr{IP: dstIP, Port: dstPort}
		key := dstAddr.String()

		mu.Lock()
		session, ok := sessions[key]
		if !ok {
			// Создаём новый PacketConn через агент
			ctx := context.Background()
			log.Printf("[UDP] creating ListenPacket for %s", dstAddr.String())
			pc, err := agent.ListenPacket(ctx, dstAddr.String())
			if err != nil {
				mu.Unlock()
				continue
			}
			log.Printf("[UDP] session created for %s", dstAddr.String())
			session = &udpSession{
				srcAddr:   src,
				dstAddr:   dstAddr,
				packetConn: pc,
				lastUsed:  time.Now(),
			}
			sessions[key] = session

			// Горутина для чтения ответов от агента и отправки клиенту
			go func(pc net.PacketConn, src *net.UDPAddr, dstIP net.IP, dstPort int) {
				respBuf := make([]byte, 65535)
				for {
					rn, _, err := pc.ReadFrom(respBuf)
					if err != nil {
						return
					}
					// Отправляем клиенту SOCKS5 UDP response
					resp := make([]byte, 3+1+4+2+rn)
					resp[0] = 0x00
					resp[1] = 0x00
					resp[2] = 0x00
					resp[3] = 0x01
					copy(resp[4:8], dstIP.To4())
					binary.BigEndian.PutUint16(resp[8:10], uint16(dstPort))
					copy(resp[10:], respBuf[:rn])
					conn.WriteToUDP(resp, src)
				}
			}(pc, src, dstIP, dstPort)
		}
		session.lastUsed = time.Now()
		mu.Unlock()

		// Отправляем данные через PacketConn (агент)
		log.Printf("[UDP] sending %d bytes to %s via agent", len(data), dstAddr.String())
		_, err = session.packetConn.WriteTo(data, dstAddr)
		if err != nil {
			mu.Lock()
			delete(sessions, key)
			mu.Unlock()
		}
	}
}
