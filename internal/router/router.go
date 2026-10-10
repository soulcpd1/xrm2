package router

import (
	"net/netip"
	"strings"

	"github.com/soulcpd1/xrm2/internal/traffic"
	"github.com/soulcpd1/xrm2/internal/xrayconfig"
)

type Engine struct {
	rules        []Rule
	exactDomains map[string]Action
	ipCache      map[string]Action
}

func New() *Engine {

	loaded, err := xrayconfig.Load("testdata/xray/config.json")
	adsRules, adsErr := xrayconfig.LoadDomainList("testdata/geosite/ads.txt")

	exactDomains := make(map[string]Action, 100000)
	cidrRules := make([]Rule, 0, 256)

	// Хелпер: сохраняет домен И с точкой, И без
	addDomain := func(domain string, action Action) {
		domain = strings.ToLower(strings.TrimSpace(domain))
		if domain == "" {
			return
		}
		exactDomains[domain] = action
		exactDomains["."+domain] = action
	}

	// 1. DefaultRules
	for _, r := range DefaultRules {
		if r.Type == DomainRule {
			addDomain(strings.TrimPrefix(r.Value, "domain:"), r.Action)
		} else {
			cidrRules = append(cidrRules, r)
		}
	}

	// 2. Реклама
	if adsErr == nil {
		for _, r := range adsRules {
			addDomain(r.Value, Block)
		}
	}

	// 3. Xray config
	if err == nil {
		for _, r := range loaded {
			action := VPS
			if r.Action == "direct" {
				action = Direct
			} else if r.Action == "block" {
				action = Block
			}

			switch r.Type {
			case xrayconfig.DomainRule:
				addDomain(strings.TrimPrefix(r.Value, "domain:"), action)
			case xrayconfig.CIDRRule:
				cidrRules = append(cidrRules, Rule{
					Type:   CIDRRule,
					Value:  r.Value,
					Action: action,
				})
			}
		}
	}

	// 4. ExtraRules + LocalRules
	for _, r := range ExtraRules {
		if r.Type == DomainRule {
			addDomain(strings.TrimPrefix(r.Value, "domain:"), r.Action)
		} else {
			cidrRules = append(cidrRules, r)
		}
	}
	for _, r := range LocalRules {
		if r.Type == DomainRule {
			addDomain(strings.TrimPrefix(r.Value, "domain:"), r.Action)
		} else {
			cidrRules = append(cidrRules, r)
		}
	}

	return &Engine{
		rules:        cidrRules,
		exactDomains: exactDomains,
		ipCache:      make(map[string]Action, 4096),
	}
}

func (e *Engine) Resolve(req *traffic.Request) Action {

	domain := strings.ToLower(req.Domain)
	if domain != "" {
		domain = strings.TrimSuffix(domain, ".")

		// 1. Точное совпадение
		if action, ok := e.exactDomains[domain]; ok {
			return action
		}

		// 2. Суффиксное совпадение (теперь работает, т.к. есть ".domain.com")
		parts := strings.Split(domain, ".")
		for i := 1; i < len(parts); i++ {
			suffix := "." + strings.Join(parts[i:], ".")
			if action, ok := e.exactDomains[suffix]; ok {
				return action
			}
		}
	}

	// 3. CIDR
	ipString := req.IP
	if ipString == "" {
		if _, err := netip.ParseAddr(domain); err == nil {
			ipString = domain
		}
	}

	if ipString != "" {
		if action, ok := e.ipCache[ipString]; ok {
			return action
		}

		if addr, err := netip.ParseAddr(ipString); err == nil {
			for _, rule := range e.rules {
				prefix, err := netip.ParsePrefix(rule.Value)
				if err == nil && prefix.Contains(addr) {
					e.ipCache[ipString] = rule.Action
					return rule.Action
				}
			}
		}
	}

	return Direct
}
