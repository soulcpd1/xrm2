package hycore

import (
	"net"
)

type udpConn struct {
	pc net.PacketConn
}

func newUDPConn(pc net.PacketConn) *udpConn {
	return &udpConn{
		pc: pc,
	}
}

func (u *udpConn) ReadFrom(b []byte) (int, string, error) {
	n, addr, err := u.pc.ReadFrom(b)
	if err != nil {
		return 0, "", err
	}
	return n, addr.String(), nil
}

func (u *udpConn) WriteTo(b []byte, addr string) (int, error) {
	ra, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return 0, err
	}

	return u.pc.WriteTo(b, ra)
}

func (u *udpConn) Close() error {
	return u.pc.Close()
}
