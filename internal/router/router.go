package router

import "github.com/soulcpd1/xrm2/internal/traffic"

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

// ONLY routing logic here (NO enums)
func (e *Engine) Resolve(req *traffic.Request) Action {

	if req.IP == "127.0.0.1" {
		return Direct
	}

	return VPS
}
