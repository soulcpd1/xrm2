package core

import (
	"fmt"
	"net"

	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Forwarder struct {
	ob *outbound.Manager
}

func NewForwarder(ob *outbound.Manager) *Forwarder {
	return &Forwarder{
		ob: ob,
	}
}

func (f *Forwarder) Dial(
	decision Decision,
	req *traffic.Request,
) (net.Conn, error) {

	tag := string(decision.Route)

	fmt.Printf(
		"[FORWARDER] action=%s host=%s\n",
		tag,
		req.Address,
	)

	conn, err := f.ob.Dial(
		tag,
		req.Network,
		req.Address,
	)

	if err != nil {
		fmt.Printf("[FORWARDER] ERROR: %v\n", err)
	}

	return conn, err
}
