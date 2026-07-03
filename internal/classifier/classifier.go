package classifier

import (
	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/rules"
)

type Classifier struct {
	engine *rules.Engine
}

func New() *Classifier {
	return &Classifier{
		engine: rules.New(),
	}
}

func (c *Classifier) Engine() *rules.Engine {
	return c.engine
}

func (c *Classifier) Resolve(req *common.Request) {
	req.Category = c.engine.Match(req.Address)
}
