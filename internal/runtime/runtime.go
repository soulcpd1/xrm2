package runtime

import (
	"github.com/soulcpd1/xrm2/internal/core"
)

type Runtime struct {
	core *core.Core
}

func New() *Runtime {
	return &Runtime{
		core: core.New(),
	}
}

func (r *Runtime) Run() error {
	_ = r.core
	return nil
}
