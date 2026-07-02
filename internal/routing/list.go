package routing

func (r *Router) Has(name string) bool {
	_, ok := r.engines[name]
	return ok
}

func (r *Router) Count() int {
	return len(r.engines)
}
