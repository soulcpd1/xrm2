package core

import "testing"

func TestCore(t *testing.T) {

	c := New()

	if c == nil {
		t.Fatal("nil")
	}

	if c.Dispatcher == nil {
		t.Fatal("dispatcher")
	}

	if c.Outbound == nil {
		t.Fatal("outbound")
	}

	if c.Routes == nil {
		t.Fatal("routes")
	}

	if c.Policy == nil {
		t.Fatal("policy")
	}

	if c.Sessions == nil {
		t.Fatal("sessions")
	}
}
