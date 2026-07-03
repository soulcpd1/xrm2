package dns

import (
	"context"
	"net"
)

type SystemResolver struct{}

func NewSystem() *SystemResolver {
	return &SystemResolver{}
}

func (r *SystemResolver) LookupIP(
	ctx context.Context,
	host string,
) ([]net.IP, error) {
	return net.DefaultResolver.LookupIP(ctx, "ip", host)
}
