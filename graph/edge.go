package graph

// Edge represents a weighted connection from one vertex to another.
type Edge struct {
	// The origin vertex id
	from string

	// The destination vertex id
	to string

	// The edge weight value
	weight float64
}

// newEdge creates a new edge initialized with given origin id, destination id,
// and weight value.
func newEdge(from, to string, weight float64) *Edge {
	return &Edge{from, to, weight}
}

// From returns the origin vertex ID of the edge.
func (e Edge) From() string {
	return e.from
}

// To returns the destination vertex ID of the edge.
func (e Edge) To() string {
	return e.to
}

// Weight returns weight value of the edge.
func (e Edge) Weight() float64 {
	return e.weight
}

// setWeight sets the weight of the current edge and returns the edge.
func (e *Edge) setWeight(v float64) *Edge {
	e.weight = v
	return e
}
