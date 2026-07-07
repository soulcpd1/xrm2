package registry

import ()

func Default() *routing.Router {

	r := routing.New()

	r.Register(tuic.NewEngine())

	return r
}
