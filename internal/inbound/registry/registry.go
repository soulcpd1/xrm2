package registry

import (
	"fmt"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/inbound/anyconnect"
	"github.com/soulcpd1/xrm2/internal/inbound/hysteria"
	"github.com/soulcpd1/xrm2/internal/inbound/tuic"
)

type Registry struct {
	dispatcher *core.Dispatcher

	hysteria   *hysteria.Handler
	tuic       *tuic.Handler
	anyconnect *anyconnect.Handler
}

func New(d *core.Dispatcher) *Registry {
	return &Registry{
		dispatcher: d,
	}
}

func (r *Registry) Init() {
	r.hysteria = hysteria.New(r.dispatcher)
	r.tuic = tuic.New(r.dispatcher)
	r.anyconnect = anyconnect.New(r.dispatcher)
}

func (r *Registry) Start() error {
	fmt.Println("[registry] inbound layer initialized")
	return nil
}

func (r *Registry) Hysteria() *hysteria.Handler {
	return r.hysteria
}

func (r *Registry) TUIC() *tuic.Handler {
	return r.tuic
}

func (r *Registry) AnyConnect() *anyconnect.Handler {
	return r.anyconnect
}
