package rules

import (
	"github.com/soulcpd1/xrm2/internal/policy"
)

func Default() *Set {

	private, _ := NewCIDRRule(
		"100.64.0.0/10",
		policy.RU,
	)

	return &Set{
		Rules: []Rule{

			DomainRule{
				Domain: "supercell.com",
				Class:  policy.Default,
			},

			SuffixRule{
				Suffix: ".ru",
				Class:  policy.RU,
			},

			KeywordRule{
				Keyword: "google",
				Class:   policy.Default,
			},

			PrivateRule{
				Class: policy.RU,
			},

			private,
		},
	}
}
