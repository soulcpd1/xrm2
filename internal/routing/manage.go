package routing

import "fmt"

func (r *Router) Start(name string) error {

	e, ok := r.Engine(name)
	if !ok {
		return fmt.Errorf("engine %s not found", name)
	}

	return e.Start()
}

func (r *Router) Stop(name string) error {

	e, ok := r.Engine(name)
	if !ok {
		return fmt.Errorf("engine %s not found", name)
	}

	return e.Stop()
}

func (r *Router) Restart(name string) error {

	e, ok := r.Engine(name)
	if !ok {
		return fmt.Errorf("engine %s not found", name)
	}

	return e.Restart()
}
