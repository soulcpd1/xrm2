package core

import (
	"fmt"

	"github.com/soulcpd1/xrm2/internal/router"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

func TestFlow(rt *router.Engine) {

	r := &traffic.Request{
		IP:      "1.1.1.1",
		Network: "tcp",
		Address: "example.com:443",
	}

	decision := rt.Resolve(r)

	fmt.Println("DECISION:", decision)
}
