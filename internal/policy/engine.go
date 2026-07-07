package policy

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/router"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (p *Engine) Resolve(req *traffic.Request) router.Action {

	// SAFETY: ensure fields exist
	domain := strings.ToLower(req.Domain)
	ip := req.IP

	// DIRECT RULES FIRST (IMPORTANT)
	if ip == "127.0.0.1" || strings.HasPrefix(ip, "192.168.") {
		return router.Direct
	}

	if strings.HasSuffix(domain, ".ru") {
		return router.Direct
	}

	// DEFAULT
	return router.VPS
}
