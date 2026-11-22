package agent

import (
	"encoding/binary"
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/soulcpd1/xrm2/internal/proto"
)

type Agent struct {
	Listen []string
}

func New(addrs []string) *Agent {
	return &Agent{
		Listen: addrs,
	}
}

func (a *Agent) Run(ctx context.Context) error {
	for _, addr := range a.Listen {
		addr := addr
		go func() {
			for {
				ln, err := net.Listen("tcp", addr)
				if err != nil {
					log.Printf("agent: cannot listen on %s: %v (retrying in 30s)", addr, err)
					select {
					case <-ctx.Done():
						return
					case <-time.After(30 * time.Second):
						continue
					}
				}
				log.Printf("agent: listening on %s", addr)
				go func() {
					<-ctx.Done()
					_ = ln.Close()
				}()
				for {
					conn, err := ln.Accept()
					if err != nil {
						select {
						case <-ctx.Done():
							return
						default:
						}
						break
					}
					go a.handle(conn)
				}
				select {
				case <-ctx.Done():
					return
				case <-time.After(5 * time.Second):
				}
			}
		}()
	}
	<-ctx.Done()
	return nil
}

func (a *Agent) handle(conn net.Conn) {
	defer conn.Close()
	// log.Printf("agent: handle from %s", conn.RemoteAddr())

	// Контекст с таймаутом 30 секунд (гарантирует завершение handle)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := proto.ReadRequest(conn)
	if err != nil {
		// log.Printf("agent: read request error: %v", err)
		return
	}
	// log.Printf("agent: request %s %s", req.Network, req.Address)

	if req.Network == "udp" {
		udpAddr, err := net.ResolveUDPAddr("udp", req.Address)
		if err != nil {
			// log.Printf("agent: resolve udp error: %v", err)
			return
		}
		udpConn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			// log.Printf("agent: dial udp error: %v", err)
			return
		}
		defer udpConn.Close()

		done := make(chan struct{})

		// Читаем из TCP (инкапсулированные пакеты с длиной) и отправляем в UDP
		go func() {
			defer close(done)
			buf := make([]byte, 65535)
			for {
				// Проверяем контекст
				select {
				case <-ctx.Done():
					return
				default:
				}
				// Устанавливаем дедлайн для чтения TCP
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))
				var length uint16
				if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
					return
				}
				if length == 0 {
					continue
				}
				if _, err := io.ReadFull(conn, buf[:length]); err != nil {
					return
				}
				// Устанавливаем дедлайн для записи UDP
				udpConn.SetWriteDeadline(time.Now().Add(5 * time.Second))
				if _, err := udpConn.Write(buf[:length]); err != nil {
					return
				}
			}
		}()

		// Читаем из UDP и отправляем в TCP (с длиной)
		buf := make([]byte, 65535)
		for {
			// Проверяем контекст
			select {
			case <-ctx.Done():
				break
			default:
			}
			// Устанавливаем дедлайн для чтения UDP
			udpConn.SetReadDeadline(time.Now().Add(5 * time.Second))
			n, err := udpConn.Read(buf)
			if err != nil {
				break
			}
			// Устанавливаем дедлайн для записи в TCP
			conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			if err := binary.Write(conn, binary.BigEndian, uint16(n)); err != nil {
				break
			}
			if _, err := conn.Write(buf[:n]); err != nil {
				break
			}
		}
		// ПРИНУДИТЕЛЬНО ЗАКРЫВАЕМ СОЕДИНЕНИЯ, ЧТОБЫ РАЗБЛОКИРОВАТЬ ГОРУТИНУ
		conn.Close()
		udpConn.Close()
		<-done
		return
	}

	// TCP
	dst, err := net.Dial(req.Network, req.Address)
	if err != nil {
		// log.Printf("agent: dial error: %v", err)
		return
	}
	defer dst.Close()
	// log.Printf("agent: connected to %s", req.Address)

	// Устанавливаем дедлайн для обоих соединений
	conn.SetDeadline(time.Now().Add(5 * time.Minute))
	dst.SetDeadline(time.Now().Add(5 * time.Minute))

	done := make(chan struct{})
	go func() {
		defer close(done)
		io.Copy(dst, conn)
		if tcp, ok := dst.(*net.TCPConn); ok {
			tcp.CloseWrite()
		}
	}()
	io.Copy(conn, dst)
	// ПРИНУДИТЕЛЬНО ЗАКРЫВАЕМ СОЕДИНЕНИЯ
	conn.Close()
	dst.Close()
	<-done
}
