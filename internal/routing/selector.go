package routing

import "fmt"

func (r *Router) Select(name string) error {

	_, ok := r.Engine(name)
	if !ok {
		return fmt.Errorf("engine %s not registered", name)
	}

	return nil
}
