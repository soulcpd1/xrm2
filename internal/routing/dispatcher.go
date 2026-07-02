package routing

import (
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/session"
)

type Dispatcher struct {
	Sessions *session.Manager
	Policy   *policy.Engine
	Routes   *route.Manager
}

func NewDispatcher(
	s *session.Manager,
	p *policy.Engine,
	r *route.Manager,
) *Dispatcher {

	return &Dispatcher{
		Sessions: s,
		Policy:   p,
		Routes:   r,
	}
}
