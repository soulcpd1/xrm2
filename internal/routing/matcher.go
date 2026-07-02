package routing

import (
	"net"
	"strings"
)

type MatchResult struct {
	Outbound string
	Reason   string
}

func (t *Table) Match(host string) MatchResult {

	host = strings.ToLower(host)

	if strings.HasSuffix(host, ".ru") {
		return MatchResult{
			Outbound: "direct",
			Reason:   "ru-domain",
		}
	}

	if ip := net.ParseIP(host); ip != nil {

		if ip.IsPrivate() {
			return MatchResult{
				Outbound: "direct",
				Reason:   "private-ip",
			}
		}
	}

	return MatchResult{
		Outbound: "nebula",
		Reason:   "default",
	}
}
