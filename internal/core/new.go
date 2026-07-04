package core

import (
	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/dns"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/rules"
	"github.com/soulcpd1/xrm2/internal/session"
)

func New(cls *classifier.Classifier) *Core {

	sessions := session.New()

	engine := policy.NewEngine()

	resolver := dns.NewSystem()

	routes := route.New([]route.Route{
		{
			Name:     "home",
			Outbound: "direct",
		},
		{
			Name:     "italy",
			Outbound: "agent",
		},
	})

	ob := outbound.Default()

	dispatcher := routing.NewDispatcher(
		sessions,
		engine,
		routes,
	)

	forwarder := NewForwarder(
		cls,
		dispatcher,
		ob,
	)

	stream := NewStreamBridge(forwarder)

	set, err := rules.Load("configs/rules.yaml")
	if err != nil {
		panic(err)
	}
	cls.Engine().Replace(set.Rules)

	return &Core{
		Sessions:   sessions,
		Policy:     engine,
		Routes:     routes,
		Dispatcher: dispatcher,
		Classifier: cls,
		DNS:        resolver,
		Outbound:   ob,
		Forwarder:  forwarder,
		Stream:     stream,
	}
}
