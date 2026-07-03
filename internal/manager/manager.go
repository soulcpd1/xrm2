package manager

import (
	"github.com/soulcpd1/xrm2/internal/registry"
	"github.com/soulcpd1/xrm2/internal/routing"
)

type Manager struct {
	router *routing.Router
}

func New() *Manager {
	return &Manager{
		router: registry.Default(),
	}
}

func (m *Manager) Router() *routing.Router {
	return m.router
}
