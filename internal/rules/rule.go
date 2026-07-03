package rules

import "github.com/soulcpd1/xrm2/internal/policy"

type Rule interface {
	Match(addr string) (policy.Category, bool)
}
