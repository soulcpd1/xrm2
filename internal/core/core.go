package core

import (
	"github.com/soulcpd1/xrm2/internal/classifier"
	"github.com/soulcpd1/xrm2/internal/outbound"
	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/session"
)

type Core struct {
	Sessions   *session.Manager
	Policy     *policy.Engine
	Routes     *route.Manager
	Outbound   *outbound.Manager
	Dispatcher *routing.Dispatcher
	Classifier *classifier.Classifier

	Stream *StreamBridge
}
