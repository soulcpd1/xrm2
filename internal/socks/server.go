package socks

import (
	"log"
	"net"

	"github.com/soulcpd1/xrm2/internal/core"
)

type Server struct {
	Dispatcher *core.Dispatcher
}

func New(d *core.Dispatcher) *Server {
	return &Server{
		Dispatcher: d,
	}
}

func (s *Server) Serve(addr string) {

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[SOCKS] listening on", addr)

	for {

		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go s.handle(conn)
	}
}
