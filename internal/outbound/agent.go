package outbound

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/soulcpd1/xrm2/internal/proto"
)

type Agent struct {
	address string
	socks5  string // если задан — подключаемся к address через этот SOCKS5
}

func NewAgent(addr string) *Agent {
	return &Agent{
		address: addr,
		socks5:  os.Getenv("XRM_AGENT_SOCKS5"),
	}
}

func (a *Agent) Name() string {
	return "agent"
}

func (a *Agent) Up() error {
	return nil
}

func (a *Agent) Down() error {
	return nil
}

func (a *Agent) Ready() bool {
	return true
}

// dialDirect — прямое TCP-подключение
func (a *Agent) dialDirect(ctx context.Context) (net.Conn, error) {
	var d net.Dialer
	return d.DialContext(ctx, "tcp", a.address)
}

// dialViaSOCKS5 — TCP до socks5-прокси, затем SOCKS5 CONNECT до a.address
func (a *Agent) dialViaSOCKS5(ctx context.Context) (net.Conn, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", a.socks5)
	if err != nil {
		return nil, fmt.Errorf("socks5 dial %s: %w", a.socks5, err)
	}

	host, portStr, err := net.SplitHostPort(a.address)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("split host/port: %w", err)
	}
	port, err := net.LookupPort("tcp", portStr)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("lookup port: %w", err)
	}

	// SOCKS5 handshake (no auth)
	if _, err := conn.Write([]byte{0x05, 0x01, 0x00}); err != nil {
		conn.Close()
		return nil, err
	}
	resp := make([]byte, 2)
	if _, err := io.ReadFull(conn, resp); err != nil {
		conn.Close()
		return nil, err
	}
	if resp[0] != 0x05 || resp[1] != 0x00 {
		conn.Close()
		return nil, fmt.Errorf("socks5 handshake failed: %v", resp)
	}

	// SOCKS5 CONNECT
	ip := net.ParseIP(host)
	var req []byte
	if ip == nil {
		// домен
		if len(host) > 255 {
			conn.Close()
			return nil, errors.New("host too long")
		}
		req = append(req, 0x05, 0x01, 0x00, 0x03, byte(len(host)))
		req = append(req, []byte(host)...)
	} else if ip4 := ip.To4(); ip4 != nil {
		// IPv4
		req = append(req, 0x05, 0x01, 0x00, 0x01)
		req = append(req, ip4...)
	} else {
		// IPv6
		req = append(req, 0x05, 0x01, 0x00, 0x04)
		req = append(req, ip.To16()...)
	}
	req = binary.BigEndian.AppendUint16(req, uint16(port))

	if _, err := conn.Write(req); err != nil {
		conn.Close()
		return nil, err
	}

	// Ответ SOCKS5
	head := make([]byte, 4)
	if _, err := io.ReadFull(conn, head); err != nil {
		conn.Close()
		return nil, err
	}
	if head[1] != 0x00 {
		conn.Close()
		return nil, fmt.Errorf("socks5 connect failed: code=%d", head[1])
	}
	// Дочитываем адрес ответа
	var addrLen int
	switch head[3] {
	case 0x01:
		addrLen = 4
	case 0x04:
		addrLen = 16
	case 0x03:
		l := make([]byte, 1)
		if _, err := io.ReadFull(conn, l); err != nil {
			conn.Close()
			return nil, err
		}
		addrLen = int(l[0])
	default:
		conn.Close()
		return nil, fmt.Errorf("socks5 bad atyp: %d", head[3])
	}
	if _, err := io.ReadFull(conn, make([]byte, addrLen+2)); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func (a *Agent) Dial(
	ctx context.Context,
	network string,
	address string,
) (net.Conn, error) {

	var conn net.Conn
	var err error

	if a.socks5 != "" {
		conn, err = a.dialViaSOCKS5(ctx)
	} else {
		conn, err = a.dialDirect(ctx)
	}
	if err != nil {
		return nil, err
	}

	err = proto.WriteRequest(conn, proto.Request{
		Network: network,
		Address: address,
	})
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func (a *Agent) ListenPacket(
	ctx context.Context,
	address string,
) (net.PacketConn, error) {

	return nil, errors.New("agent UDP not implemented yet")
}
