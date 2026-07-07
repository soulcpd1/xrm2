package common

import "net"

// Request is unified traffic envelope for ALL protocols
type Request struct {
	Network string
	Address string
	Domain  string
	IP      string

	// Protocol source (hysteria / tuic / anyconnect / future)
	Protocol string

	// Raw connection (IMPORTANT: used by forwarder)
	Conn net.Conn
}
