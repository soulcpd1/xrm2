package hysteria

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
)

type Handler struct {
	forwarder *core.Forwarder
}

func New(f *core.Forwarder) *Handler {
	return &Handler{forwarder: f}
}

// REAL ENTRYPOINT
func (h *Handler) Serve(conn net.Conn) {

	req := core.ToRequest(conn)

	h.forwarder.Dial(req)
}
