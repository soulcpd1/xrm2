package routing

import (
	"time"

	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/session"
)

func (d *Dispatcher) Resolve(
	sessionID string,
	category policy.Category,
) (*route.Route, error) {

	if s, ok := d.Sessions.Get(sessionID); ok {
		d.Sessions.Touch(sessionID)
		return s.Route, nil
	}

	decision := d.Policy.Resolve(category)

	rt, err := d.Routes.Get(string(decision.Route))
	if err != nil {
		return nil, err
	}

	d.Sessions.Put(sessionID, &session.Session{
		ID:        sessionID,
		Route:     rt,
		CreatedAt: time.Now(),
		LastSeen:  time.Now(),
	})

	return rt, nil
}
