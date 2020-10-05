package graph

type edgeMap map[string]*edge

func (m edgeMap) has(id string) bool {
	_, exists := m[id]
	return exists
}

func (m edgeMap) get(id string) *edge {
	if !m.has(id) {
		return nil
	}
	return m[id]
}

func (m edgeMap) add(id string, e edge) bool {
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

// A vertex represent a node in a graph. It consists of a unique id
// with associated data, and a collection of edges connecting it
// to other vertices.
type vertex struct {
	// id is a unique identifier in the scope of the graph.
	// It should never be modified.
	id string

	// the value of the vertex. It can be of any type.
	value interface{}

	// edges is a collection of edges stored in a map for quick access
	// and insert/remove operations.
	edges edgeMap
}

func newVertex(id string, v interface{}) *vertex {
	return &vertex{id, v, edgeMap{}}
}

func (v vertex) Len() int {
	return len(v.edges)
}

func (v vertex) Edges() []edge {
	values := []edge{}
	for _, e := range v.edges {
		values = append(values, *e)
	}
	return values
}

func (v *vertex) addEdge(e *edge) bool {
	return v.edges.add(e.to, *e)
}

func (v *vertex) hasEdge(id string) bool {
	return v.edges.has(id)
}

func (v *vertex) getEdge(id string) *edge {
	return v.edges.get(id)
}

func (v *vertex) removeEdge(id string) bool {
	return v.edges.remove(id)
}

func (v vertex) walkEdges(f func(e edge)) {
	for _, e := range v.edges {
		f(*e)
	}
}
