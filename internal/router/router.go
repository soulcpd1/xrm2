package router

import (
	"strings"

	"github.com/soulcpd1/xrm2/internal/traffic"
	"github.com/soulcpd1/xrm2/internal/xrayconfig"
)

type Engine struct {
	rules []Rule
}

func New() *Engine {

	loaded, err := xrayconfig.Load("testdata/xray/config.json")

	rules := DefaultRules

	if err == nil {

		rules = make([]Rule, 0, len(loaded))

		for _, r := range loaded {

			action := VPS
			if r.Action == "direct" {
				action = Direct
			}

			switch r.Type {

			case xrayconfig.DomainRule:
				rules = append(rules, Rule{
					Type:   DomainRule,
					Value:  r.Value,
					Action: action,
				})

			case xrayconfig.CIDRRule:
				rules = append(rules, Rule{
					Type:   CIDRRule,
					Value:  r.Value,
					Action: action,
				})
			}
		}
	}

	return &Engine{
		rules: rules,
	}
}

func (e *Engine) Resolve(req *traffic.Request) Action {

	domain := strings.ToLower(req.Domain)
	ip := req.IP

	for _, rule := range e.rules {

		switch rule.Type {

		case CIDRRule:

			if strings.HasPrefix(ip, "127.") &&
				rule.Value == "127.0.0.0/8" {
				return rule.Action
			}

		case DomainRule:

			value := strings.TrimPrefix(rule.Value, "domain:")

			if strings.HasSuffix(domain, value) {
				return rule.Action
			}
		}
	}

	return Direct
}
