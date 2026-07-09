package socks

import (
	"io"
	"log"
	"net"
)

func (s *Server) handle(client net.Conn) {

	defer client.Close()

	if err := s.handshake(client); err != nil {
		log.Println("[SOCKS] handshake:", err)
		return
	}

	cmd, req, err := s.parseRequest(client)
	if err != nil {
		log.Println("[SOCKS] request:", err)
		return
	}

	switch cmd {

	case 0x03:

		log.Println("[SOCKS] UDP ASSOCIATE")

		if err := s.handleUDPAssociate(client); err != nil {
			log.Println("[SOCKS] udp:", err)
		}

		return

	case 0x01:

		log.Printf(
			"[SOCKS] CONNECT network=%s host=%s address=%s",
			req.Network,
			req.Domain,
			req.Address,
		)

	default:

		log.Printf("[SOCKS] unknown cmd=0x%02x", cmd)
		return
	}

	remote, err := s.Dispatcher.Handle(req)
	if err != nil {
		log.Println("[SOCKS] dial:", err)
		return
	}
	defer remote.Close()

	_, err = client.Write([]byte{
		0x05,
		0x00,
		0x00,
		0x01,
		0, 0, 0, 0,
		0, 0,
	})
	if err != nil {
		return
	}

	log.Println("[SOCKS] relay started")

	go io.Copy(remote, client)
	io.Copy(client, remote)

	log.Println("[SOCKS] relay finished")
}
