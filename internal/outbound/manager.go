package outbound

import "fmt"

type Manager struct {
	drivers map[string]Driver
}

func New() *Manager {
	return &Manager{
		drivers: make(map[string]Driver),
	}
}

func (m *Manager) Register(d Driver) {
	m.drivers[d.Name()] = d
}

func (m *Manager) Driver(name string) (Driver, error) {
	d, ok := m.drivers[name]
	if !ok {
		return nil, fmt.Errorf("driver %s not found", name)
	}
	return d, nil
}

func (m *Manager) Names() []string {
	names := make([]string, 0, len(m.drivers))
	for name := range m.drivers {
		names = append(names, name)
	}
	return names
}
