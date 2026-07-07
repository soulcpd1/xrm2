package socks

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/soulcpd1/xrm2/internal/traffic"
)

func (s *Server) parseRequest(conn net.Conn) (*traffic.Request, error) {

	var hdr [4]byte

	if _, err := io.ReadFull(conn, hdr[:]); err != nil {
		return nil, err
	}

	// VER CMD RSV ATYP
	if hdr[0] != 0x05 {
		return nil, fmt.Errorf("invalid socks version")
	}

	// only CONNECT
	if hdr[1] != 0x01 {
		return nil, fmt.Errorf("unsupported command")
	}

	var host string

	switch hdr[3] {

	case 0x01:

		var ip [4]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return nil, err
		}

		host = net.IP(ip[:]).String()

	case 0x03:

		var l [1]byte

		if _, err := io.ReadFull(conn, l[:]); err != nil {
			return nil, err
		}

		buf := make([]byte, int(l[0]))

		if _, err := io.ReadFull(conn, buf); err != nil {
			return nil, err
		}

		host = string(buf)

	case 0x04:

		var ip [16]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return nil, err
		}

		host = net.IP(ip[:]).String()

	default:

		return nil, fmt.Errorf("unsupported atyp")
	}

	var portbuf [2]byte

	if _, err := io.ReadFull(conn, portbuf[:]); err != nil {
		return nil, err
	}

	port := binary.BigEndian.Uint16(portbuf[:])

	req := &traffic.Request{
		Network: "tcp",
		Domain:  host,
		Address: fmt.Sprintf("%s:%d", host, port),
	}

	return req, nil
}
