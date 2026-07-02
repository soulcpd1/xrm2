package routing

func (r *Router) Names() []string {

	names := make([]string, 0, len(r.engines))

	for name := range r.engines {
		names = append(names, name)
	}

	return names
}
