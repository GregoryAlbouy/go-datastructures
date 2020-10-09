package graph

type edgeMap map[string]*Edge

func (m edgeMap) has(id string) bool {
	_, exists := m[id]
	return exists
}

func (m edgeMap) get(id string) *Edge {
	if !m.has(id) {
		return nil
	}
	return m[id]
}

func (m edgeMap) add(id string, e Edge) bool {
	if m.has(id) {
		return false
	}
	m[id] = &e
	return true
}

func (m edgeMap) remove(id string) bool {
	if !m.has(id) {
		return false
	}
	delete(m, id)
	return true
}
