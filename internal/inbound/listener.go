package inbound

import (
	"log"
	"net"

	"github.com/soulcpd1/xrm2/internal/runtime"
)

type Server struct {
	rt *runtime.Runtime
}

func New(rt *runtime.Runtime) *Server {
	return &Server{rt: rt}
}

func (s *Server) Listen(addr string) error {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("[INBOUND] listening on", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("[INBOUND] accept error:", err)
			continue
		}

		go s.rt.Handle(conn)
	}
}
