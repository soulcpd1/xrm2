package engine

type Engine interface {
	Name() string

	Start() error
	Stop() error
	Restart() error
}
