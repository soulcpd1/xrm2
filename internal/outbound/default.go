package outbound

func Default() *Manager {

	m := New()

	m.Register(NewDirect())
	m.Register(NewNebula())
	m.Register(NewAgent("192.168.100.1:7000"))

	return m
}
