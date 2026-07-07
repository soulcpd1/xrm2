package anyconnect

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
	"github.com/soulcpd1/xrm2/internal/traffic"
)

// Handler is adapter for ocserv (Cisco AnyConnect compatible server)
type Handler struct {
	F *core.Forwarder
}

func New(f *core.Forwarder) *Handler {
	return &Handler{F: f}
}

// Handle is called after TLS tunnel is established by ocserv
func (h *Handler) Handle(conn net.Conn, network string, address string) error {

	req := &traffic.Request{
		Network:  network,
		Address:  address,
		Domain:   address,
		IP:       conn.RemoteAddr().String(),
		Protocol: "anyconnect",
		Conn:     conn,
	}

	// CRITICAL: unified pipeline entry
	_, err := h.F.Dial(req)
	return err
}
