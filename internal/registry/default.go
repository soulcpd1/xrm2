package registry

import (
	"github.com/soulcpd1/xrm2/internal/hysteria"
	"github.com/soulcpd1/xrm2/internal/routing"
	"github.com/soulcpd1/xrm2/internal/tuic"
)

func Default() *routing.Router {

	r := routing.New()

	r.Register(hysteria.NewEngine())
	r.Register(tuic.NewEngine())

	return r
}
