package hycore

import (
	"context"

	hyserver "github.com/apernet/hysteria/core/v2/server"
	"github.com/google/uuid"

	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/policy"
)

func (o *Outbound) UDP(reqAddr string) (hyserver.UDPConn, error) {

	req := &common.Request{
		SessionID: uuid.NewString(),
		Network:   "udp",
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

	pc, err := o.core.Forwarder.ListenPacket(
		context.Background(),
		driver,
		":0",
	)
	if err != nil {
		return nil, err
	}

	return newUDPConn(pc), nil
}
