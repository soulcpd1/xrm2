package router

type RuleType uint8

const (
	DomainRule RuleType = iota
	CIDRRule
)

type Rule struct {
	Type   RuleType
	Value  string
	Action Action
}

var DefaultRules = []Rule{
	{
		Type:   CIDRRule,
		Value:  "127.0.0.0/8",
		Action: Direct,
	},
}
