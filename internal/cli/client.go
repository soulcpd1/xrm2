package cli

import (
	"fmt"
	"github.com/soulcpd1/xrm2/internal/core"
)

func Run() {
	engine := core.NewEngine()

	req := core.Request{
		Address: "rebus.rustore.ru",
	}

	dec := engine.Resolve(req)

	fmt.Println("ROUTE:", dec.Route, dec.Reason)
}
