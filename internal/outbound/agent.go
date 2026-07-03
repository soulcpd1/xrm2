package outbound

import (
	"context"
	"errors"
	"net"
)

type Agent struct {
	address string
}

func NewAgent(addr string) *Agent {
	return &Agent{
		address: addr,
	}
}

func (a *Agent) Name() string {
	return "agent"
}

func (a *Agent) Up() error {
	return nil
}

func (a *Agent) Down() error {
	return nil
}

func (a *Agent) Ready() bool {
	return true
}

func (a *Agent) Dial(
	ctx context.Context,
	network string,
	address string,
) (net.Conn, error) {

	var d net.Dialer

	// Пока соединяемся только с Agent.
	// Передача proto.Request будет следующим этапом.
	return d.DialContext(ctx, "tcp", a.address)
}

func (a *Agent) ListenPacket(
	ctx context.Context,
	address string,
) (net.PacketConn, error) {

	return nil, errors.New("agent UDP not implemented yet")
}
