package socks

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/soulcpd1/xrm2/internal/traffic"
)

func (s *Server) parseRequest(conn net.Conn) (byte, *traffic.Request, error) {

	var hdr [4]byte

	if _, err := io.ReadFull(conn, hdr[:]); err != nil {
		return 0, nil, err
	}

	if hdr[0] != 0x05 {
		return 0, nil, fmt.Errorf("invalid socks version")
	}

	cmd := hdr[1]

	if cmd == 0x03 {
		return cmd, &traffic.Request{
			Network: "udp",
		}, nil
	}

	if cmd != 0x01 {
		return 0, nil, fmt.Errorf(
			"unsupported command: cmd=0x%02x atyp=0x%02x",
			hdr[1],
			hdr[3],
		)
	}

	var host string

	switch hdr[3] {

	case 0x01:

		var ip [4]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return 0, nil, err
		}

		host = net.IP(ip[:]).String()

	case 0x03:

		var l [1]byte

		if _, err := io.ReadFull(conn, l[:]); err != nil {
			return 0, nil, err
		}

		buf := make([]byte, int(l[0]))

		if _, err := io.ReadFull(conn, buf); err != nil {
			return 0, nil, err
		}

		host = string(buf)

	case 0x04:

		var ip [16]byte

		if _, err := io.ReadFull(conn, ip[:]); err != nil {
			return 0, nil, err
		}

		host = net.IP(ip[:]).String()

	default:
		return 0, nil, fmt.Errorf("unsupported atyp")
	}

	var portbuf [2]byte

	if _, err := io.ReadFull(conn, portbuf[:]); err != nil {
		return 0, nil, err
	}

	port := binary.BigEndian.Uint16(portbuf[:])

	return cmd, &traffic.Request{
		Network: "tcp",
		Domain:  host,
		Address: fmt.Sprintf("%s:%d", host, port),
	}, nil
}
