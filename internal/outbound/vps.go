package outbound

import (
	"context"
	"net"
	"os"
	"time"
)

const defaultAgentAddress = "192.168.100.1:7000"

type VPSDriver struct {
	agent *Agent
}

func NewVPSDriver() VPSDriver {

	addr := os.Getenv("XRM_AGENT")

	if addr == "" {
		addr = defaultAgentAddress
	}

	return VPSDriver{
		agent: NewAgent(addr),
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
