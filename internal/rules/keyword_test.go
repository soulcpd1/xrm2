package rules

import (
	"testing"

	"github.com/soulcpd1/xrm2/internal/policy"
)

func TestKeywordRule(t *testing.T) {

	r := KeywordRule{
		Keyword: "google",
		Class:   policy.RU,
	}

	tests := []struct {
		addr string
		want bool
	}{
		{"google.com", true},
		{"www.google.com", true},
		{"dns.google", true},
		{"youtube.com", false},
		{"1.1.1.1", false},
	}

	for _, tt := range tests {

		_, ok := r.Match(tt.addr)

		if ok != tt.want {
			t.Fatalf("%s => %v", tt.addr, ok)
		}
	}
}
