package rules

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/policy"
)

func TestCIDRRule(t *testing.T) {

	r, err := NewCIDRRule(
		"192.168.0.0/16",
		policy.RU,
	)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		addr string
		want bool
	}{
		{"192.168.1.1", true},
		{"192.168.100.10:443", true},
		{"8.8.8.8", false},
		{"google.com", false},
	}

	for _, tt := range tests {

		_, ok := r.Match(tt.addr)

		if ok != tt.want {
			t.Fatalf("%s => %v", tt.addr, ok)
		}
	}
}
