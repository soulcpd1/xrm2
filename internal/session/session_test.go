package session

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/route"
)

func TestSessionManager(t *testing.T) {

	m := New()

	m.Put("abc", &Session{
		ID:    "abc",
		Route: &route.Route{Name: "italy"},
	})

	if m.Count() != 1 {
		t.Fatal("count != 1")
	}

	s, ok := m.Get("abc")
	if !ok {
		t.Fatal("session not found")
	}

	if s.Route.Name != "italy" {
		t.Fatal("wrong route")
	}

	m.Delete("abc")

	if m.Count() != 0 {
		t.Fatal("count != 0")
	}
}
