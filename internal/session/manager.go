package session

func (m *Manager) Get(id string) (*Session, bool) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.sessions[id]
	return s, ok
}

func (m *Manager) Put(id string, s *Session) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[id] = s
}

func (m *Manager) Delete(id string) {

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, id)
}

func (m *Manager) Count() int {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.sessions)
}
