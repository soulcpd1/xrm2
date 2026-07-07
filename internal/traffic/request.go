package traffic

import "net"

type Request struct {
	Network  string
	Address  string
	Domain   string
	IP       string
	Protocol string
	Conn     net.Conn
}
