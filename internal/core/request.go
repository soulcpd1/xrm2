package core

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/traffic"
)

func ToRequest(conn net.Conn) *traffic.Request {
	addr := conn.RemoteAddr().String()

	return &traffic.Request{
		Network:  "tcp",
		Address:  addr,
		IP:       addr,
		Domain:   "",
		Protocol: "hysteria",
		Conn:     conn,
	}
}
