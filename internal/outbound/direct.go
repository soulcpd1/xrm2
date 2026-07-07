package outbound

import (
	"net"
)

type DirectDriver struct{}

func NewDirectDriver() DirectDriver {
	return DirectDriver{}
}

func (d DirectDriver) Dial(network, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}
