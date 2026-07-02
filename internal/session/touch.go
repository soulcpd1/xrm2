package session

import "time"

func (m *Manager) Touch(id string) {

	m.mu.Lock()
	defer m.mu.Unlock()

	if s, ok := m.sessions[id]; ok {
		s.LastSeen = time.Now()
	}
}
