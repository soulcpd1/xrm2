package rules

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/soulcpd1/xrm2/internal/policy"
)

func LoadFile(path string) (*Engine, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	e := New()

	for _, r := range cfg.Rules {
		cat := policy.Category(r.Policy)

		switch r.Type {

		case "domain":
			e.Add(DomainRule{
				Domain: r.Value,
				Class:  cat,
			})

		case "suffix":
			e.Add(SuffixRule{
				Suffix: r.Value,
				Class:  cat,
			})

		case "keyword":
			e.Add(KeywordRule{
				Keyword: r.Value,
				Class:   cat,
			})

		case "cidr":
			e.Add(NewCIDRRule(
				r.Value,
				cat,
			))

		case "private":
			e.Add(PrivateRule{
				Class: cat,
			})
		}
	}

	return e, nil
}
