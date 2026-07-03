package common

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/policy"
	"github.com/soulcpd1/xrm2/internal/route"
)

type Request struct {
	SessionID string

	Conn net.Conn

	Network string

	Host string
	Port string

	Address string

	Category policy.Category

	Route *route.Route
}
