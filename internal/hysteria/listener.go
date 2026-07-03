package hysteria

import (
	"net"

	"github.com/soulcpd1/xrm2/internal/session"
)

type Listener struct {
	handler *Handler
}

func NewListener(h *Handler) *Listener {
	return &Listener{handler: h}
}

func (l *Listener) Serve(addr string) error {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {

			req := BuildRequest(c)

			req.SessionID = session.NewID()

			_ = l.handler.Resolve(req)

		}(conn)
	}
}
