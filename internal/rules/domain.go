package rules

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type DomainRule struct {
	Domain string
	Class  policy.Category
}

func (r DomainRule) Match(addr string) (policy.Category, bool) {
	if strings.Contains(addr, r.Domain) {
		return r.Class, true
	}
	return "", false
}
