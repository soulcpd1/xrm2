package rules

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type CIDRRule struct {
	Network *net.IPNet
	Class   policy.Category
}

func NewCIDRRule(cidr string, class policy.Category) (*CIDRRule, error) {
	_, n, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &CIDRRule{
		Network: n,
		Class:   class,
	}, nil
}

func (r *CIDRRule) Match(addr string) (policy.Category, bool) {

	host := addr

	if h, _, err := net.SplitHostPort(addr); err == nil {
		host = h
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return "", false
	}

	if r.Network.Contains(ip) {
		return r.Class, true
	}

	return "", false
}
