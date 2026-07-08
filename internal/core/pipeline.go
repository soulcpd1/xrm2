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

	if req.IP == "" && req.Domain != "" {

		ips, err := net.LookupIP(req.Domain)
		if err == nil && len(ips) > 0 {
			req.IP = ips[0].String()
		}
	}

	action := p.router.Resolve(req)

	decision := Decision{
		Route:  Route(action),
		Reason: "router",
	}

	return p.forwarder.Dial(decision, req)
}
