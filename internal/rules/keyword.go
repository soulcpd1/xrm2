package rules

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type KeywordRule struct {
	Keyword string
	Class   policy.Category
}

func (r KeywordRule) Match(addr string) (policy.Category, bool) {

	if strings.Contains(
		strings.ToLower(addr),
		strings.ToLower(r.Keyword),
	) {
		return r.Class, true
	}

	return "", false
}
