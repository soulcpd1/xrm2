package routing

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/session"
)

func TestDispatcher(t *testing.T) {

	sessions := session.New()

	engine := policy.NewEngine()

	routes := route.New(policy.DefaultRoutes())

	d := NewDispatcher(
		sessions,
		engine,
		routes,
	)

	rt, err := d.Resolve(
		"google.com",
		policy.Blocked,
	)

	if err != nil {
		t.Fatal(err)
	}

	if rt.Name != "italy" {
		t.Fatalf("expected italy, got %s", rt.Name)
	}
}
