package policy

import "github.com/soulcpd1/xrm2/internal/route"

type Rule struct {
	Category Category
	Route    route.Name
}

type Engine struct {
	rules []Rule
}

func NewEngine() *Engine {
	return &Engine{
		rules: []Rule{
			{
				Category: RU,
				Route:    route.Home,
			},
			{
				Category: Private,
				Route:    route.Home,
			},
			{
				Category: Blocked,
				Route:    route.Italy,
			},
			{
				Category: Default,
				Route:    route.Italy,
			},
		},
	}
}

func (e *Engine) Resolve(cat Category) Decision {

	for _, r := range e.rules {
		if r.Category == cat {
			return Decision{
				Route:  r.Route,
				Reason: string(cat),
			}
		}
	}

	return Decision{
		Route:  route.Italy,
		Reason: "fallback",
	}
}
