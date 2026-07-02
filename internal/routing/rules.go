package routing

type Rule struct {
	Name   string
	Engine string

	Domain []string
	IPCIDR []string

	Default bool
}

type Rules []Rule

func (r Rules) DefaultRule() *Rule {

	for i := range r {
		if r[i].Default {
			return &r[i]
		}
	}

	return nil
}
