package core

import (
	"context"
	"net"

	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/routing"
)

type Forwarder struct {
	Classifier *classifier.Classifier
	Dispatcher *routing.Dispatcher
	Outbound   *outbound.Manager
}

func NewForwarder(
	cls *classifier.Classifier,
	disp *routing.Dispatcher,
	ob *outbound.Manager,
) *Forwarder {
	return &Forwarder{
		Classifier: cls,
		Dispatcher: disp,
		Outbound:   ob,
	}
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

func (f *Forwarder) ListenPacket(
	ctx context.Context,
	driver string,
	address string,
) (net.PacketConn, error) {

	d, err := f.Outbound.Driver(driver)
	if err != nil {
		return nil, err
	}

	return d.ListenPacket(ctx, address)
}
