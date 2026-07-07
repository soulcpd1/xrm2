package core

import "context"

type Route string

const (
	Direct Route = "direct"
	Proxy  Route = "proxy"
	Home   Route = "home"
)

type Request struct {
	Address string
}

type Decision struct {
	Route  Route
	Reason string
}

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Resolve(req Request) Decision {
	switch {
	case req.Address == "rebus.rustore.ru":
		return Decision{Route: Home, Reason: "ru"}
	default:
		return Decision{Route: Proxy, Reason: "default"}
	}
}

func (e *Engine) Handle(ctx context.Context, req Request) Decision {
	return e.Resolve(req)
}
