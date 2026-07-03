package rules

import "github.com/soulcpd1/xrm2/internal/policy"

type Engine struct {
	rules []Rule
}

func New() *Engine {
	return &Engine{
		rules: make([]Rule, 0),
	}
}

func (e *Engine) Add(r Rule) {
	e.rules = append(e.rules, r)
}

func (e *Engine) Match(addr string) policy.Category {

	for _, r := range e.rules {
		if cat, ok := r.Match(addr); ok {
			return cat
		}
	}

	return policy.Default
}
