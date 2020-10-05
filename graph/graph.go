package graph

// Work in progress

type adjacencyList map[string]*vertex

type graph struct {
	data          adjacencyList
	isDirected    bool
	defaultWeight int
}

// Interface of a graph
type Interface interface {
	SetDirected(b bool) Interface
	SetDefaultWeight(v int) Interface
	Clear() Interface

	Add(id string, value interface{}) bool
	AddMany(vertices ...struct {
		id    string
		value interface{}
	})
	Has(id string) bool
	Get(id string) *vertex
	Remove(id string) bool
	ResetVertex(id string) bool

	GetEdge(from, to string) *edge
	AddEdge(from, to string, weight ...int) bool
	RemoveEdge(from, to string) bool
	GetEdgeWeight(from, to string) int
	SetEdgeWeight(from, to string, weight int) bool

	// debug()
}

// New returns a graph
func New() Interface {
	return (&graph{}).
		SetDirected(false).
		SetDefaultWeight(1).
		Clear()
}

func (g *graph) SetDirected(b bool) Interface {
	g.isDirected = b
	return g
}

func (g *graph) SetDefaultWeight(v int) Interface {
	g.defaultWeight = v
	return g
}

func (g *graph) Clear() Interface {
	g.data = adjacencyList{}
	return g
}

func (g *graph) Add(id string, v interface{}) bool {
	if g.Has(id) {
		return false
	}
	g.data[id] = newVertex(id, v)
	return true
}

func (g *graph) AddMany(vertices ...struct {
	id    string
	value interface{}
}) {
	for _, v := range vertices {
		g.Add(v.id, v.value)
	}
}

func (g graph) Has(id string) bool {
	_, exists := g.data[id]
	return exists
}

func (g graph) Get(id string) *vertex {
	if !g.Has(id) {
		return nil
	}
	return g.data[id]
}

func (g *graph) Remove(id string) bool {
	if !g.Has(id) {
		return false
	}
	delete(g.data, id)
	for _, v := range g.data {
		v.removeEdge(id)
	}
	return true
}

func (g *graph) ResetVertex(id string) bool {
	if v := g.Get(id); v != nil {
		return g.Remove(v.id) && g.Add(v.id, v.value)
	}
	return false
}

func (g *graph) AddEdge(from, to string, weight ...int) bool {
	w := g.defaultWeight
	if len(weight) > 0 {
		w = weight[0]
	}

	src, dst, isSafeAdd, _ := g.checkEdgeOps(from, to)

	if !isSafeAdd {
		return false
	}

	srcAdded := src.addEdge(newEdge(from, to, w))
	if g.isDirected {
		return srcAdded
	}

	dstAdded := dst.addEdge(newEdge(to, from, w))
	return srcAdded && dstAdded
}

func (g *graph) RemoveEdge(from, to string) bool {
	src, dst, _, isSafeRem := g.checkEdgeOps(from, to)

	if !isSafeRem {
		return false
	}

	srcRemoved := src.removeEdge(to)
	if g.isDirected {
		return srcRemoved
	}

	dstRemoved := dst.removeEdge(from)
	return srcRemoved && dstRemoved
}

func (g graph) GetEdge(from, to string) *edge {
	v := g.Get(from)
	if v == nil {
		return nil
	}
	return v.getEdge(to)
}

func (g *graph) SetEdgeWeight(from, to string, weight int) bool {
	if e := g.GetEdge(from, to); e != nil {
		e.weight = weight
		return true
	}
	return false
}

func (g graph) GetEdgeWeight(from, to string) int {
	if e := g.GetEdge(from, to); e != nil {
		return e.weight
	}
	return 0
}

func (g graph) checkEdgeOps(from, to string) (*vertex, *vertex, bool, bool) {
	src := g.Get(from)
	dst := g.Get(to)
	bothExist := src != nil && dst != nil
	isSafeAdd := bothExist &&
		!src.hasEdge(to) &&
		(!dst.hasEdge(from) || g.isDirected)
	isSafeRem := bothExist &&
		src.hasEdge(to) &&
		(dst.hasEdge(from) || g.isDirected)

	return src, dst, isSafeAdd, isSafeRem
}
