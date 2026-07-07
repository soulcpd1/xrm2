package socks

import (
	"fmt"
	"io"
	"net"
)

func (s *Server) handshake(conn net.Conn) error {

	var hdr [2]byte

	if _, err := io.ReadFull(conn, hdr[:]); err != nil {
		return err
	}

	if hdr[0] != 0x05 {
		return fmt.Errorf("not socks5")
	}

	methods := make([]byte, int(hdr[1]))

	if _, err := io.ReadFull(conn, methods); err != nil {
		return err
	}

	_, err := conn.Write([]byte{
		0x05,
		0x00,
	})

	return err
}
