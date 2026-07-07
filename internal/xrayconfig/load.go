package xrayconfig

import (
	"encoding/json"
	"os"
)

type RuleType uint8

const (
	DomainRule RuleType = iota
	CIDRRule
)

type Rule struct {
	Type   RuleType
	Value  string
	Action string
}

type Config struct {
	Routing Routing `json:"routing"`
}

type Routing struct {
	Rules []XRayRule `json:"rules"`
}

type XRayRule struct {
	Domain      []string `json:"domain"`
	IP          []string `json:"ip"`
	OutboundTag string   `json:"outboundTag"`
}

func Load(path string) ([]Rule, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	out := make([]Rule, 0)

	for _, r := range cfg.Routing.Rules {

		action := "vps"
		if r.OutboundTag == "direct" {
			action = "direct"
		}

		for _, d := range r.Domain {
			out = append(out, Rule{
				Type:   DomainRule,
				Value:  d,
				Action: action,
			})
		}

		for _, ip := range r.IP {
			out = append(out, Rule{
				Type:   CIDRRule,
				Value:  ip,
				Action: action,
			})
		}
	}

	return out, nil
}
