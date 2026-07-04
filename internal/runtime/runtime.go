package runtime

import (
	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/hycore"
	"github.com/soulcpd1/xrm2/internal/hysteria"
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
	engine := hysteria.NewEngine()

	cfg, err := engine.Load()
	if err != nil {
		return err
	}

	serverCfg, err := hycore.BuildConfig(cfg)
	if err != nil {
		return err
	}

	serverCfg.Outbound = hycore.NewOutbound(r.core)

	srv, err := hycore.New(serverCfg)
	if err != nil {
		return err
	}

	return srv.Run()
}
