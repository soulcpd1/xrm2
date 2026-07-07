package registry

import (
	"fmt"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/inbound/anyconnect"
	"github.com/soulcpd1/xrm2/internal/inbound/hysteria"
	"github.com/soulcpd1/xrm2/internal/inbound/tuic"
)

// Registry is control-plane for all inbound protocols
type Registry struct {
	forwarder *core.Forwarder

	hysteria   *hysteria.Handler
	tuic       *tuic.Handler
	anyconnect *anyconnect.Handler
}

func New(f *core.Forwarder) *Registry {
	return &Registry{
		forwarder: f,
	}
}

func (r *Registry) Init() {
	r.hysteria = hysteria.New(r.forwarder)
	r.tuic = tuic.New(r.forwarder)
	r.anyconnect = anyconnect.New(r.forwarder)
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
