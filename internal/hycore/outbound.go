package hycore

import (
	"context"
	"net"

	hyserver "github.com/apernet/hysteria/core/v2/server"

	"github.com/google/uuid"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/policy"
)

type Outbound struct {
	core *core.Core
}

func NewOutbound(c *core.Core) hyserver.Outbound {
	return &Outbound{
		core: c,
	}
}

func (o *Outbound) TCP(reqAddr string) (net.Conn, error) {

	route, err := o.core.Dispatcher.Resolve(
		uuid.NewString(),
		policy.Default,
	)
	if err != nil {
		return nil, err
	}

	return o.core.Stream.Dial(
		context.Background(),
		route.Outbound,
		"tcp",
		reqAddr,
	)
}

func (o *Outbound) UDP(reqAddr string) (hyserver.UDPConn, error) {
	return nil, nil
}

func (o *Outbound) CheckUDP(reqAddr string) error {
	return nil
}
