package route

import "fmt"

type Manager struct {
	routes []Route
}

func New(routes []Route) *Manager {
	return &Manager{
		routes: routes,
	}
}

func (m *Manager) Get(name string) (*Route, error) {

	for i := range m.routes {
		if m.routes[i].Name == name {
			return &m.routes[i], nil
		}
	}

	return nil, fmt.Errorf("route %s not found", name)
}

func (m *Manager) All() []Route {
	return m.routes
}
