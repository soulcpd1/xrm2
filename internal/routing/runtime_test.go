package routing

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/session"
)

func TestRuntimeRouting(t *testing.T) {

	routes := route.New([]route.Route{
		{
			Name:     "home",
			Outbound: "direct",
		},
		{
			Name:     "italy",
			Outbound: "agent",
		},
	})

	engine := policy.NewEngine()

	dispatcher := NewDispatcher(
		session.New(),
		engine,
		routes,
	)

	tests := []struct {
		name     string
		category policy.Category
		want     string
	}{
		{"RU", policy.RU, "direct"},
		{"DEFAULT", policy.Default, "agent"},
		{"BLOCKED", policy.Blocked, "agent"},
	}

	for _, tt := range tests {

		req := &common.Request{
			Category: tt.category,
		}

		if err := dispatcher.Resolve(req); err != nil {
			t.Fatal(err)
		}

		if req.Route.Outbound != tt.want {
			t.Fatalf("%s: got %s want %s",
				tt.name,
				req.Route.Outbound,
				tt.want,
			)
		}
	}
}
