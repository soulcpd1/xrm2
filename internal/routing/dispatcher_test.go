package routing

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/common"
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

	req := &common.Request{
		Address:  "google.com",
		Category: policy.Blocked,
	}

	if err := d.Resolve(req); err != nil {
		t.Fatal(err)
	}

	if req.Route == nil {
		t.Fatal("route is nil")
	}

	if req.Route.Name != "italy" {
		t.Fatalf("expected italy, got %s", req.Route.Name)
	}
}
