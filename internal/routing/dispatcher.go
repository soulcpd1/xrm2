package routing

import (
	"github.com/soulcpd1/xrm2/internal/common"
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

func (d *Dispatcher) Resolve(req *common.Request) error {
	decision := d.Policy.Resolve(req.Category)

	rt, err := d.Routes.Get(string(decision.Route))
	if err != nil {
		return err
	}

	req.Route = rt
	return nil
}
