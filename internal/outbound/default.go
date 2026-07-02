package outbound

func Default() *Manager {

	m := New()

	m.Register(NewDirect())
	m.Register(NewNebula())

	return m
}
