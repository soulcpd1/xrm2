package router

import (
	"net/netip"
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

		rules = make([]Rule, 0, len(loaded)+len(ExtraRules))

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

		rules = append(rules, ExtraRules...)
	}

	return &Engine{
		rules: rules,
	}
}

func (e *Engine) Resolve(req *traffic.Request) Action {

	domain := strings.ToLower(req.Domain)

	ipString := req.IP

	if ipString == "" {
		if _, err := netip.ParseAddr(domain); err == nil {
			ipString = domain
		}
	}

	var (
		ip    netip.Addr
		hasIP bool
	)

	if ipString != "" {

		addr, err := netip.ParseAddr(ipString)

		if err == nil {
			ip = addr
			hasIP = true
		}
	}

	for _, rule := range e.rules {

		switch rule.Type {

		case CIDRRule:

			if hasIP {

				prefix, err := netip.ParsePrefix(rule.Value)

				if err == nil && prefix.Contains(ip) {
					return rule.Action
				}
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
