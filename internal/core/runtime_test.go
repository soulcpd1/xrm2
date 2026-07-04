package core

import (
	"context"
	"testing"

	"github.com/soulcpd1/xrm2/internal/common"
	"github.com/soulcpd1/xrm2/internal/policy"
)

func TestPipelineRouting(t *testing.T) {

	c := New()

	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "RU",
			addr: "ya.ru:443",
			want: "direct",
		},
		{
			name: "WORLD",
			addr: "google.com:443",
			want: "agent",
		},
	}

	for _, tt := range tests {

		req := &common.Request{
			Network: "tcp",
			Address: tt.addr,
		}

		c.Classifier.Resolve(req)

		if err := c.Dispatcher.Resolve(req); err != nil {
			t.Fatal(err)
		}

		driver, err := c.Forwarder.SelectDriver(req)
		if err != nil {
			t.Fatal(err)
		}

		if driver.Name() != tt.want {
			t.Fatalf("%s: got %s want %s",
				tt.name,
				driver.Name(),
				tt.want,
			)
		}

		if req.Category == policy.Default &&
			tt.want == "direct" {
			t.Fatal("unexpected routing")
		}
	}
}
