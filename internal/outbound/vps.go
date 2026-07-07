package outbound

import (
	"context"
	"net"
	"time"
)

type VPSDriver struct {
	agent *Agent
}

func NewVPSDriver() VPSDriver {
	return VPSDriver{
		agent: NewAgent("192.168.100.1:7000"),
	}
}

func (v VPSDriver) Dial(network, addr string) (net.Conn, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	return v.agent.Dial(
		ctx,
		network,
		addr,
	)
}
