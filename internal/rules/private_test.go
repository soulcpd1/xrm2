package rules

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/policy"
)

func TestPrivateRule(t *testing.T) {

	r := PrivateRule{
		Class: policy.RU,
	}

	tests := []struct {
		addr string
		want bool
	}{
		{"127.0.0.1", true},
		{"10.0.0.1", true},
		{"172.16.1.1", true},
		{"192.168.1.10", true},
		{"169.254.1.1", true},
		{"8.8.8.8", false},
		{"1.1.1.1", false},
		{"google.com", false},
	}

	for _, tt := range tests {

		_, ok := r.Match(tt.addr)

		if ok != tt.want {
			t.Fatalf("%s => %v", tt.addr, ok)
		}
	}
}
