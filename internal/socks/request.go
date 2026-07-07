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

	if hdr[0] != 0x05 {
		return nil, fmt.Errorf("invalid socks version")
	}

	if hdr[1] != 0x01 {
		return nil, fmt.Errorf("unsupported command")
	}

	var (
		host     string
		domain   string
		ipString string
	)

	switch hdr[3] {

	case 0x01:

		var ip [4]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return nil, err
		}

		ipString = net.IP(ip[:]).String()
		host = ipString

	case 0x03:

		var l [1]byte

		if _, err := io.ReadFull(conn, l[:]); err != nil {
			return nil, err
		}

		buf := make([]byte, int(l[0]))

		if _, err := io.ReadFull(conn, buf); err != nil {
			return nil, err
		}

		domain = string(buf)
		host = domain

	case 0x04:

		var ip [16]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return nil, err
		}

		ipString = net.IP(ip[:]).String()
		host = ipString

	default:
		return nil, fmt.Errorf("unsupported atyp")
	}

	var portbuf [2]byte

	if _, err := io.ReadFull(conn, portbuf[:]); err != nil {
		return nil, err
	}

	port := binary.BigEndian.Uint16(portbuf[:])

	return &traffic.Request{
		Network: "tcp",
		Domain:  domain,
		IP:      ipString,
		Address: fmt.Sprintf("%s:%d", host, port),
	}, nil
}
