package runtime

import "fmt"

type Runtime struct{}

func New() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Run() error {
	return fmt.Errorf("runtime not implemented")
}
