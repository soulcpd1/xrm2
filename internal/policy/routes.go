package policy

import "github.com/soulcpd1/xrm2/internal/route"

func DefaultRoutes() []route.Route {

	return []route.Route{

		{
			Name:     "home",
			Outbound: "home",
			Priority: 10,
		},

		{
			Name:     "italy",
			Outbound: "italy",
			Priority: 100,
			Fallback: "home",
		},
	}
}
