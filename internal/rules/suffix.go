package rules

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type SuffixRule struct {
	Suffix string
	Class  policy.Category
}

func (r SuffixRule) Match(addr string) (policy.Category, bool) {
	if strings.HasSuffix(addr, r.Suffix) {
		return r.Class, true
	}
	return "", false
}
