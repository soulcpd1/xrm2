package routing

type Table struct {
	Outbounds []Outbound
}

func DefaultTable() *Table {
	return &Table{
		Outbounds: []Outbound{
			{
				Name:      "direct",
				Type:      Direct,
				Interface: "",
				Enabled:   true,
			},
			{
				Name:      "nebula",
				Type:      Nebula,
				Interface: "nebula1",
				Enabled:   true,
			},
		},
	}
}
