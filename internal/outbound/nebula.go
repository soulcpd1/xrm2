package outbound

import (
	"context"
	"net"
)

type Nebula struct{}

func NewNebula() *Nebula {
	return &Nebula{}
}

func (n *Nebula) Name() string {
	return "nebula"
}

func (n *Nebula) Up() error {
	return nil
}

func (n *Nebula) Down() error {
	return nil
}

func (n *Nebula) Ready() bool {

	iface, err := net.InterfaceByName("nebula1")
	if err != nil {
		return false
	}

	return iface.Flags&net.FlagUp != 0
}

func (n *Nebula) Dial(
	ctx context.Context,
	network string,
	address string,
) (net.Conn, error) {

	var nd net.Dialer

	return nd.DialContext(
		ctx,
		network,
		address,
	)
}
