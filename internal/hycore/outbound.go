package hycore

import (
	"context"
	"net"

	hyserver "github.com/apernet/hysteria/core/v2/server"

	"github.com/soulcpd1/xrm2/internal/core"
)

type Outbound struct {
	core *core.Core
}

func NewOutbound(c *core.Core) hyserver.Outbound {
	return &Outbound{core: c}
}

func (o *Outbound) TCP(reqAddr string) (net.Conn, error) {
	return o.core.Forwarder.Dial(
		context.Background(),
		"tcp",
		reqAddr,
	)
}

func (o *Outbound) CheckUDP(reqAddr string) error {
	return nil
}
