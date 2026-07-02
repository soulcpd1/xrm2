package session

func (m *Manager) Cleanup() int {

	m.mu.Lock()
	defer m.mu.Unlock()

	removed := 0

	for id, s := range m.sessions {

		if s.Expired() {
			delete(m.sessions, id)
			removed++
		}
	}

	return removed
}
