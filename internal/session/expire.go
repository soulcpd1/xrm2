package session

import "time"

func (s *Session) Expired() bool {
	return time.Since(s.LastSeen) > DefaultTTL
}
