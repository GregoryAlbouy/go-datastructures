package graph

// edge represents a weighted connection from one vertex to another.
type edge struct {
	// The origin vertex id
	from string

	// The destination vertex id
	to string

	// The edge weight value
	weight int
}

// newEdge creates a new edge initialized with given origin id, destination id,
// and weight value.
func newEdge(from, to string, weight int) *edge {
	return &edge{from, to, weight}
}

// setWeight sets the weight of the current edge and returns the edge.
func (e *edge) setWeight(v int) *edge {
	e.weight = v
	return e
}
