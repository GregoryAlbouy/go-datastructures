package graph

// A Vertex represents a node in a graph. It consists of a unique id
// with associated data, and a collection of edges connecting it
// to other vertices.
type Vertex struct {
	// id is a unique identifier in the scope of the graph.
	// It should never be modified.
	id string

	// the value of the vertex. It can be of any type.
	value interface{}

	// edges is a collection of edges stored in a map for quick access
	// and insert/remove operations.
	edges edgeMap
}

func newVertex(id string, v interface{}) *Vertex {
	return &Vertex{id, v, edgeMap{}}
}

// ID returns the unique ID of the vertex.
func (v Vertex) ID() string {
	return v.id
}

// Value returns the value of the vertex.
func (v Vertex) Value() interface{} {
	return v.value
}

// Edges returns a slice of edges connected to the vertex.
func (v Vertex) Edges() []Edge {
	values := []Edge{}
	for _, e := range v.edges {
		values = append(values, *e)
	}
	return values
}

// Len returns the number of edges connected to the Vertex.
func (v Vertex) Len() int {
	return len(v.edges)
}

func (v *Vertex) addEdge(e *Edge) bool {
	return v.edges.add(e.to, *e)
}

func (v *Vertex) hasEdge(id string) bool {
	return v.edges.has(id)
}

func (v *Vertex) getEdge(id string) *Edge {
	return v.edges.get(id)
}

func (v *Vertex) removeEdge(id string) bool {
	return v.edges.remove(id)
}

func (v Vertex) walkEdges(f func(e Edge)) {
	for _, e := range v.edges {
		f(*e)
	}
}
