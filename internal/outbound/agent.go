package outbound

import (
	"context"
	"errors"
	"net"

	"github.com/soulcpd1/xrm2/internal/proto"
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

	conn, err := d.DialContext(ctx, "tcp", a.address)
	if err != nil {
		return nil, err
	}

	err = proto.WriteRequest(conn, proto.Request{
		Network: network,
		Address: address,
	})
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func (a *Agent) ListenPacket(
	ctx context.Context,
	address string,
) (net.PacketConn, error) {

	return nil, errors.New("agent UDP not implemented yet")
}
