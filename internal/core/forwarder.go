package core

import (
	"fmt"
	"net"

	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Forwarder struct {
	policy *policy.Engine
	ob     *outbound.Manager
}

func NewForwarder(p *policy.Engine, ob *outbound.Manager) *Forwarder {
	return &Forwarder{policy: p, ob: ob}
}

func (f *Forwarder) Dial(req *traffic.Request) (net.Conn, error) {

	action := f.policy.Resolve(req)

	fmt.Printf("[FORWARDER] action=%s host=%s\n", action, req.Address)

	conn, err := f.ob.Dial(string(action), req.Network, req.Address)

	if err != nil {
		fmt.Printf("[FORWARDER] ERROR: %v\n", err)
	}

	return conn, err
}
