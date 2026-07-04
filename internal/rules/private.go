package rules

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type PrivateRule struct {
	Class policy.Category
}

func (r PrivateRule) Match(addr string) (policy.Category, bool) {

	host := addr

	if h, _, err := net.SplitHostPort(addr); err == nil {
		host = h
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return "", false
	}

	if ip.IsPrivate() ||
		ip.IsLoopback() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() {

		return r.Class, true
	}

	return "", false
}
