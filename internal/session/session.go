package session

import (
	"sync"
	"time"

	"github.com/soulcpd1/xrm2/internal/route"
)

type Session struct {
	ID        string
	Route     *route.Route
	CreatedAt time.Time
	LastSeen  time.Time
}

type Manager struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

func New() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}
