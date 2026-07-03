package common

import "net"

type Request struct {
	SessionID string

	Conn net.Conn

	Network string
	Address string
}
