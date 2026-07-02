package routing

import "github.com/soulcpd1/xrm2/internal/engine"

type Router struct {
	engines map[string]engine.Engine
}

func New() *Router {
	return &Router{
		engines: make(map[string]engine.Engine),
	}
}

func (r *Router) Register(e engine.Engine) {
	r.engines[e.Name()] = e
}

func (r *Router) Engine(name string) (engine.Engine, bool) {
	e, ok := r.engines[name]
	return e, ok
}

func (r *Router) Engines() []engine.Engine {

	list := make([]engine.Engine, 0, len(r.engines))

	for _, e := range r.engines {
		list = append(list, e)
	}

	return list
}
