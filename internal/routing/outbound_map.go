package routing

type Exit struct {
	Name   string
	Driver string
}

func DefaultExits() []Exit {
	return []Exit{
		{
			Name:   "home",
			Driver: "direct",
		},
		{
			Name:   "italy",
			Driver: "nebula",
		},
	}
}
