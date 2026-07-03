package agent

import (
	"context"
	"net"
)

type Agent struct {
	Listen string
}

func New(addr string) *Agent {
	return &Agent{
		Listen: addr,
	}
}

func (a *Agent) Run(ctx context.Context) error {
	ln, err := net.Listen("tcp", a.Listen)
	if err != nil {
		return err
	}
	defer ln.Close()

	go func() {
		<-ctx.Done()
		_ = ln.Close()
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
			}
			continue
		}

		go a.handle(conn)
	}
}

func (a *Agent) handle(conn net.Conn) {
	defer conn.Close()

	// TODO:
	// read destination
	// dial target
	// bridge streams
}
