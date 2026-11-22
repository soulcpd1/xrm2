package outbound

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/soulcpd1/xrm2/internal/proto"
)

type Agent struct {
	address       string
	fallbackAgent string // второй xrm-agent (например, ZeroTier), если основной недоступен
	socks5        string // если задан — подключаемся к address через этот SOCKS5 (режим агента)
	directSocks5  string // если задан — прямой режим: xrm работает как чистый SOCKS5-клиент
}

func NewAgent(addr string) *Agent {
	return &Agent{
		address:       addr,                          // основной xrm-agent (XRM_AGENT)
		fallbackAgent: os.Getenv("XRM_AGENT_FALLBACK"), // второй xrm-agent (XRM_AGENT_FALLBACK)
		socks5:        os.Getenv("XRM_AGENT_SOCKS5"),   // fallback SOCKS5
		directSocks5:  os.Getenv("XRM_DIRECT_SOCKS5"),  // прямой SOCKS5 (без агента)
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
	d.Timeout = 2 * time.Second
	return d.DialContext(ctx, "tcp", a.address)
}

func (a *Agent) dialDirectTo(ctx context.Context, addr string) (net.Conn, error) {
	var d net.Dialer
	d.Timeout = 2 * time.Second
	return d.DialContext(ctx, "tcp", addr)
}

// dialViaSOCKS5 — TCP до socks5-прокси, затем SOCKS5 CONNECT до target
func (a *Agent) dialViaSOCKS5(ctx context.Context, target string) (net.Conn, error) {
	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", a.socks5)
	if err != nil {
		return nil, fmt.Errorf("socks5 dial %s: %w", a.socks5, err)
	}

	host, portStr, err := net.SplitHostPort(target)
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

	// Прямой режим: xrm работает как чистый SOCKS5-клиент (без агента).
	if a.directSocks5 != "" {
		return a.dialViaSOCKS5(ctx, address)
	}

	// Основной путь: прямой xrm-agent (XRM_AGENT), XRM_AGENT_SOCKS5 используется только в Fallback 2
	log.Printf("[XRM] Dial: trying main path (agent=%s)", a.address)
	conn, err := a.dialDirect(ctx)
	if err == nil {
		err = proto.WriteRequest(conn, proto.Request{
			Network: network,
			Address: address,
		})
		if err == nil {
			log.Printf("[XRM] Dial: WriteRequest sent to agent")
			return conn, nil
		}
		conn.Close()
	}
	log.Printf("[XRM] Dial: main path FAILED: %v", err)

	// Fallback 1: если основной agent недоступен, пробуем второй agent (XRM_AGENT_FALLBACK)
	if a.fallbackAgent != "" {
		log.Printf("[XRM] Dial: trying fallback agent (%s)", a.fallbackAgent)
		conn2, err2 := a.dialDirectTo(ctx, a.fallbackAgent)
		if err2 == nil {
			err2 = proto.WriteRequest(conn2, proto.Request{
				Network: network,
				Address: address,
			})
			if err2 == nil {
				log.Printf("[XRM] Dial: fallback agent OK")
				return conn2, nil
			}
			conn2.Close()
		}
		log.Printf("[XRM] Dial: fallback agent FAILED: %v", err2)
	}

	// Fallback 2: если и второй agent недоступен, пробуем SOCKS5-резерв (XRM_AGENT_SOCKS5)
	if a.socks5 != "" {
		log.Printf("[XRM] Dial: trying fallback (socks5=%s)", a.socks5)
		conn2, err2 := a.dialViaSOCKS5(ctx, address)
		if err2 == nil {
			log.Printf("[XRM] Dial: fallback OK")
			return conn2, nil
		}
		log.Printf("[XRM] Dial: fallback FAILED: %v", err2)
	}

	return nil, fmt.Errorf("agent dial failed: %w (fallback also failed)", err)
}


func (a *Agent) ListenPacket(ctx context.Context, address string) (net.PacketConn, error) {
	log.Printf("[XRM] ListenPacket called for %s", address)
	conn, err := a.Dial(ctx, "udp", address)
	if err != nil {
		return nil, err
	}
	return &udpPacketConn{conn: conn}, nil
}

type udpPacketConn struct {
	conn net.Conn
}

func (c *udpPacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	var length uint16
	if err := binary.Read(c.conn, binary.BigEndian, &length); err != nil {
		return 0, nil, err
	}
	n, err = io.ReadFull(c.conn, p[:length])
	if err != nil {
		return 0, nil, err
	}
	return n, nil, nil
}

func (c *udpPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	if err := binary.Write(c.conn, binary.BigEndian, uint16(len(p))); err != nil {
		return 0, err
	}
	return c.conn.Write(p)
}

func (c *udpPacketConn) Close() error { return c.conn.Close() }
func (c *udpPacketConn) LocalAddr() net.Addr { return c.conn.LocalAddr() }
func (c *udpPacketConn) SetDeadline(t time.Time) error { return c.conn.SetDeadline(t) }
func (c *udpPacketConn) SetReadDeadline(t time.Time) error { return c.conn.SetReadDeadline(t) }
func (c *udpPacketConn) SetWriteDeadline(t time.Time) error { return c.conn.SetWriteDeadline(t) }
