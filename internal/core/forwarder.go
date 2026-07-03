package core

import (
	"context"
	"net"

	"github.com/soulcpd1/xrm2/internal/outbound"
)

type Forwarder struct {
	Outbound *outbound.Manager
}

func NewForwarder(ob *outbound.Manager) *Forwarder {
	return &Forwarder{Outbound: ob}
}

func (f *Forwarder) Dial(
	ctx context.Context,
	driver string,
	network string,
	address string,
) (net.Conn, error) {

	d, err := f.Outbound.Driver(driver)
	if err != nil {
		return nil, err
	}

	return d.Dial(ctx, network, address)
}
