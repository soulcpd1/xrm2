package classifier

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/policy"
)

type Classifier struct{}

func New() *Classifier {
	return &Classifier{}
}

func (c *Classifier) Category(address string) policy.Category {

	host := strings.ToLower(address)

	switch {

	case strings.Contains(host, "clash"):
		return policy.Default

	case strings.Contains(host, "supercell"):
		return policy.Default

	default:
		return policy.Default
	}
}
