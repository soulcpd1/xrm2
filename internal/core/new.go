package core

import (
	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/session"
)

func New() *Core {

	sessions := session.New()
	engine := policy.NewEngine()
	routes := route.New(policy.DefaultRoutes())
	ob := outbound.New()

	cls := classifier.New()

	dispatcher := routing.NewDispatcher(
		sessions,
		engine,
		routes,
	)

	forwarder := NewForwarder(ob)
	stream := NewStreamBridge(forwarder)

	return &Core{
		Sessions:   sessions,
		Policy:     engine,
		Routes:     routes,
		Outbound:   ob,
		Dispatcher: dispatcher,
		Classifier: cls,
		Stream:     stream,
	}
}
