package classifier

import (
	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/policy"
)

func (c *Classifier) Resolve(req *common.Request) policy.Category {
	return c.Category(req.Address)
}
