package outbound

type Driver interface {
	Name() string

	Up() error
	Down() error
	Ready() bool

	Forward(any) error
}
