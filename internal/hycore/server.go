package hycore

import (
	"fmt"

	hyserver "github.com/apernet/hysteria/core/v2/server"
)

type Server struct {
	server hyserver.Server
}

func New(cfg *hyserver.Config) (*Server, error) {
	s, err := hyserver.NewServer(cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		server: s,
	}, nil
}

func (s *Server) Run() error {
	if s.server == nil {
		return fmt.Errorf("nil server")
	}

	return s.server.Serve()
}

func (s *Server) Close() error {
	if s.server == nil {
		return nil
	}

	return s.server.Close()
}
