package core

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/router"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Pipeline struct {
	router    *router.Engine
	forwarder *Forwarder
}

func NewPipeline(
	r *router.Engine,
	f *Forwarder,
) *Pipeline {

	return &Pipeline{
		router:    r,
		forwarder: f,
	}
}

func (p *Pipeline) Handle(
	req *traffic.Request,
) (net.Conn, error) {

	action := p.router.Resolve(req)

	decision := Decision{
		Route:  Route(action),
		Reason: "router",
	}

	return p.forwarder.Dial(decision, req)
}
