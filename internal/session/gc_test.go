package session

import (
	"testing"
	"time"

	"github.com/soulcpd1/xrm2/internal/route"
)

func TestCleanup(t *testing.T) {

	m := New()

	m.Put("alive", &Session{
		ID:        "alive",
		Route:     &route.Route{Name: "home"},
		CreatedAt: time.Now(),
		LastSeen:  time.Now(),
	})

	m.Put("dead", &Session{
		ID:        "dead",
		Route:     &route.Route{Name: "italy"},
		CreatedAt: time.Now().Add(-10 * time.Minute),
		LastSeen:  time.Now().Add(-24 * time.Hour),
	})

	n := m.Cleanup()

	if n != 1 {
		t.Fatalf("removed=%d", n)
	}

	if m.Count() != 1 {
		t.Fatalf("count=%d", m.Count())
	}
}
