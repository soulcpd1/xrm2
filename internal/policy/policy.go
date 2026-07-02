package policy

import "github.com/soulcpd1/xrm2/internal/route"

type Decision struct {
	Route  route.Name
	Reason string
}
