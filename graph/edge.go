package graph

// edge represents a weighted connection from one vertex to another.
type edge struct {
	// The origin vertex id
	from string

	// The destination vertex id
	to string

	// The edge weight value
	weight float64
}

// newEdge creates a new edge initialized with given origin id, destination id,
// and weight value.
func newEdge(from, to string, weight float64) *edge {
	return &edge{from, to, weight}
}

// setWeight sets the weight of the current edge and returns the edge.
func (e *edge) setWeight(v float64) *edge {
	e.weight = v
	return e
}
