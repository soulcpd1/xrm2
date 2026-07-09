package socks

import (
	"io"
	"net"
)

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

	for {

		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			return
		}

		_ = n
		_ = src
	}
}
