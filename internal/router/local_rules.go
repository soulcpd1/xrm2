package router

var LocalRules = []Rule{
	{
		Type:   DomainRule,
		Value:  "domain:max.ru",
		Action: Block,
	},
}
