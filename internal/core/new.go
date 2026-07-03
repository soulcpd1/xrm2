package core

import (
	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/rules"
	"github.com/soulcpd1/xrm2/internal/session"
)

func New() *Core {

	sessions := session.New()
	engine := policy.NewEngine()

	ob := outbound.New()

	dispatcher := routing.NewDispatcher(
		sessions,
		engine,
		nil,
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
		Dispatcher: dispatcher,
		Classifier: cls,
		Outbound:   ob,
	}
}
