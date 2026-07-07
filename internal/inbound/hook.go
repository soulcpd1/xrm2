package inbound

import (
	"fmt"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Hook struct {
	dispatcher *core.Dispatcher
}

func NewHook(d *core.Dispatcher) *Hook {
	return &Hook{dispatcher: d}
}

// 🔥 THIS IS THE REAL ENTRY POINT FOR ALL PROTOCOLS
func (h *Hook) Handle(req *traffic.Request) {

	fmt.Printf("[INBOUND] proto=%s addr=%s ip=%s domain=%s\n",
		req.Protocol, req.Address, req.IP, req.Domain)

	_, err := h.dispatcher.Handle(req)

	if err != nil {
		fmt.Printf("[INBOUND] ERROR: %v\n", err)
	}
}
