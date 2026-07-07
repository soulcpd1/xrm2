package tuic

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

type Handler struct {
	F *core.Forwarder
}

func New(f *core.Forwarder) *Handler {
	return &Handler{F: f}
}

// Entry point from TUIC server
func (h *Handler) Handle(conn net.Conn, network string, address string) error {

	req := &traffic.Request{
		Network:  network,
		Address:  address,
		Domain:   address,
		IP:       conn.RemoteAddr().String(),
		Protocol: "tuic",
		Conn:     conn,
	}

	// CRITICAL: same pipeline as hysteria
	_, err := h.F.Dial(req)
	return err
}
