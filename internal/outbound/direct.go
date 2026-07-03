package outbound

import (
	"context"
	"net"
)

type Direct struct{}

func NewDirect() *Direct {
	return &Direct{}
}

func (d *Direct) Name() string {
	return "direct"
}

func (d *Direct) Up() error {
	return nil
}

func (d *Direct) Down() error {
	return nil
}

func (d *Direct) Ready() bool {
	return true
}

func (d *Direct) Dial(
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
