package hycore

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/hysteria"
)

type Authenticator struct {
	cfg *hysteria.Auth
}

func NewAuthenticator(cfg *hysteria.Auth) *Authenticator {
	return &Authenticator{
		cfg: cfg,
	}
}

func (a *Authenticator) Authenticate(addr net.Addr, auth string, tx uint64) (bool, string) {
	if a.cfg == nil {
		return false, ""
	}

	if a.cfg.Type != "password" {
		return false, ""
	}

	if auth != a.cfg.Password {
		return false, ""
	}

	return true, "default"
}
