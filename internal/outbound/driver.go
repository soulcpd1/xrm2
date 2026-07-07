package outbound

import "net"

type Driver interface {
	Name() string
	Dial(network, addr string) (net.Conn, error)
}
