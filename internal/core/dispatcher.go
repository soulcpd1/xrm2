package core

import (
	"fmt"
	"net"

	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Dispatcher struct {
	forwarder *Forwarder
	policy    *policy.Engine
}

func NewDispatcher(f *Forwarder, p *policy.Engine) *Dispatcher {
	return &Dispatcher{
		forwarder: f,
		policy:    p,
	}
}

func (d *Dispatcher) Handle(req *traffic.Request) (net.Conn, error) {

	// 🔥 STEP 1: inbound dump
	fmt.Printf("\n[TRACE] ===== INBOUND =====\n")
	fmt.Printf("[TRACE] proto=%s\n", req.Protocol)
	fmt.Printf("[TRACE] domain=%s\n", req.Domain)
	fmt.Printf("[TRACE] ip=%s\n", req.IP)
	fmt.Printf("[TRACE] addr=%s\n", req.Address)

	// 🔥 STEP 2: policy decision
	action := d.policy.Resolve(req)

	fmt.Printf("[TRACE] POLICY RESULT = %s\n", action)

	// 🔥 STEP 3: outbound execution
	conn, err := d.forwarder.Dial(req)

	if err != nil {
		fmt.Printf("[TRACE] ERROR: %v\n", err)
	}

	fmt.Printf("[TRACE] FLOW END\n")

	return conn, err
}
