package graph

func reverseInPlace(a []*Vertex) []*Vertex {
	n := len(a)

	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}

	return a
}

func (g graph) checkEdgeOps(from, to string) (*Vertex, *Vertex, bool, bool) {
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

func dfsRecurse(g graph, id string, f func(Vertex) interface{}, done map[string]bool) {
	vx := g.Get(id)

	if vx == nil {
		return
	}

	done[id] = true
	f(*vx)

	vx.walkEdges(func(e Edge) {
		if _, isDone := done[e.to]; !isDone {
			dfsRecurse(g, id, f, done)
		}
	})
}
