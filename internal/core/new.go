package core

import (
	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/rules"
	"github.com/soulcpd1/xrm2/internal/session"
)

func New() *Core {

	sessions := session.New()
	engine := policy.NewEngine()

	routes := route.New([]route.Route{
		{
			Name:     "home",
			Outbound: "direct",
		},
		{
			Name:     "italy",
			Outbound: "nebula",
		},
	})

	ob := outbound.Default()

	forwarder := NewForwarder(ob)
	stream := NewStreamBridge(forwarder)

	dispatcher := routing.NewDispatcher(
		sessions,
		engine,
		routes,
	)

	cls := classifier.New()

	// временные правила (MVP)
	cls.Engine().Add(rules.DomainRule{
		Domain: "supercell.com",
		Class:  policy.Default,
	})

	cls.Engine().Add(rules.SuffixRule{
		Suffix: ".ru",
		Class:  policy.RU,
	})

	return &Core{
		Sessions:   sessions,
		Policy:     engine,
		Routes:     routes,
		Dispatcher: dispatcher,
		Classifier: cls,
		Outbound:   ob,
		Forwarder:  forwarder,
		Stream:     stream,
	}
}
