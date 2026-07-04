package policy

import "github.com/soulcpd1/xrm2/internal/route"

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Resolve(cat Category) Decision {
	switch cat {
	case "ru":
		return Decision{Route: route.Name("italy")}
	case "default":
		return Decision{Route: route.Name("direct")}
	default:
		return Decision{Route: route.Name("home")}
	}
}
