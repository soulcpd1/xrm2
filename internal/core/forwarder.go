package core

import (
	"context"
	"net"

	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/route"
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

func (f *Forwarder) SelectDriver(
	req *common.Request,
) (outbound.Driver, error) {

	return f.Outbound.Driver(req.Route.Outbound)
}

func (f *Forwarder) Dial(
	ctx context.Context,
	network string,
	address string,
) (net.Conn, error) {

	req := &common.Request{
		Network: network,
		Address: address,
	}

	f.Classifier.Resolve(req)

	if err := f.Dispatcher.Resolve(req); err != nil {
		return nil, err
	}

	d, err := f.Outbound.Driver(req.Route.Outbound)
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
