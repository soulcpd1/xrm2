package core

import (
	"fmt"
	"net"

	"github.com/soulcpd1/xrm2/internal/router"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Dispatcher struct {
	forwarder *Forwarder
	router    *router.Engine
}

func NewDispatcher(f *Forwarder, r *router.Engine) *Dispatcher {
	return &Dispatcher{
		forwarder: f,
		router:    r,
	}
}

func (d *Dispatcher) Handle(req *traffic.Request) (net.Conn, error) {

	fmt.Printf("\n[TRACE] ===== INBOUND =====\n")
	fmt.Printf("[TRACE] proto=%s\n", req.Protocol)
	fmt.Printf("[TRACE] domain=%s\n", req.Domain)
	fmt.Printf("[TRACE] ip=%s\n", req.IP)
	fmt.Printf("[TRACE] addr=%s\n", req.Address)

	action := d.router.Resolve(req)

	decision := Decision{
		Route:  Route(action),
		Reason: "router",
	}

	fmt.Printf("[TRACE] DECISION = %s\n", decision.Route)

	conn, err := d.forwarder.Dial(decision, req)

	if err != nil {
		fmt.Printf("[TRACE] ERROR: %v\n", err)
	}

	fmt.Printf("[TRACE] FLOW END\n")

	return conn, err
}
