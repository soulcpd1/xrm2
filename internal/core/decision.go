package core

type Route string

const (
	Direct Route = "direct"
	Proxy  Route = "proxy"
	Home   Route = "home"
)

type Decision struct {
	Route  Route
	Reason string
}
