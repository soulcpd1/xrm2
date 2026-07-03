package dns

import (
	"context"
	"net"
)

type Resolver interface {
	LookupIP(ctx context.Context, host string) ([]net.IP, error)
}
