package hysteria

import (
	"io"
	"log"
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
)

type Handler struct {
	dispatcher *core.Dispatcher
}

func New(d *core.Dispatcher) *Handler {
	return &Handler{
		dispatcher: d,
	}
}

func (h *Handler) Serve(conn net.Conn) {
	defer conn.Close()

	req := core.ToRequest(conn)

	remote, err := h.dispatcher.Handle(req)
	if err != nil {
		log.Println("[HYSTERIA] dispatch:", err)
		return
	}
	defer remote.Close()

	log.Println("[HYSTERIA] relay started")

	go io.Copy(remote, conn)
	io.Copy(conn, remote)

	log.Println("[HYSTERIA] relay finished")
}
