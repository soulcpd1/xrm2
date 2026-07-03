package outbound

import (
	"context"
	"net"
)

type Driver interface {
	Name() string

	Up() error
	Down() error
	Ready() bool

	Dial(
		ctx context.Context,
		network string,
		address string,
	) (net.Conn, error)
}
