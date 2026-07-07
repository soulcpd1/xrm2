package outbound

import (
	"fmt"
	"net"
)

type Manager struct {
	direct DirectDriver
	vps    VPSDriver
}

func NewManager(d DirectDriver, v VPSDriver) *Manager {
	return &Manager{
		direct: d,
		vps:    v,
	}
}

func (m *Manager) Dial(tag, network, addr string) (net.Conn, error) {

	switch tag {

	case "direct":
		return m.direct.Dial(network, addr)

	case "vps":
		return m.vps.Dial(network, addr)

	default:
		return nil, fmt.Errorf("unknown outbound tag: %s", tag)
	}
}
