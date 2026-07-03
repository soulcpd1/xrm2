package hycore

import (
	"context"
	"net"

	hyserver "github.com/apernet/hysteria/core/v2/server"
	"github.com/google/uuid"

	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/policy"
)

type Outbound struct {
	core *core.Core
}

func NewOutbound(c *core.Core) hyserver.Outbound {
	return &Outbound{core: c}
}

func (o *Outbound) TCP(reqAddr string) (net.Conn, error) {

	req := &common.Request{
		SessionID: uuid.NewString(),
		Network:   "tcp",
		Address:   reqAddr,
		Category:  policy.Default,
	}

	if err := o.core.Dispatcher.Resolve(req); err != nil {
		return nil, err
	}

	driver := "direct"
	if req.Route != nil {
		driver = req.Route.Outbound
	}

	return o.core.Forwarder.Dial(
		context.Background(),
		driver,
		req.Network,
		req.Address,
	)
}

func (o *Outbound) UDP(reqAddr string) (hyserver.UDPConn, error) {
	return nil, nil
}

func (o *Outbound) CheckUDP(reqAddr string) error {
	return nil
}
