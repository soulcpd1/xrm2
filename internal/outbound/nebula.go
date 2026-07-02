package outbound

import "net"

type Nebula struct{}

func NewNebula() *Nebula {
	return &Nebula{}
}

func (n *Nebula) Name() string {
	return "nebula"
}

func (n *Nebula) Up() error {
	return nil
}

func (n *Nebula) Down() error {
	return nil
}

func (n *Nebula) Ready() bool {

	iface, err := net.InterfaceByName("nebula1")
	if err != nil {
		return false
	}

	if iface.Flags&net.FlagUp == 0 {
		return false
	}

	return true
}

func (n *Nebula) Forward(conn any) error {
        return nil
}
