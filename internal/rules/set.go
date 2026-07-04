package rules

type Set struct {
	Rules []Rule
}

func (s *Set) Apply(e *Engine) {
	for _, r := range s.Rules {
		e.Add(r)
	}
}
