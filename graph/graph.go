package graph

import (
	"fmt"
	"math"

	pqueue "github.com/gregoryalbouy/go-datastructures/priorityqueue"
	"github.com/gregoryalbouy/go-datastructures/queue"
)

type adjacencyList map[string]*Vertex

// FilterFunc type
type FilterFunc func(v Vertex) bool

type graph struct {
	data          adjacencyList
	isDirected    bool
	defaultWeight float64
}

// Interface of a graph
type Interface interface {
	SetDirected(isDirected bool) Interface
	SetDefaultWeight(weight float64) Interface
	Clear() Interface

	Add(id string, value interface{}) bool
	AddMany(vertices ...struct {
		id    string
		value interface{}
	})
	Has(id string) bool
	Get(id string) *Vertex
	Remove(id string) bool
	ResetVertex(id string) bool

	GetEdge(from, to string) *Edge
	AddEdge(from, to string, weight ...float64) bool
	RemoveEdge(from, to string) bool
	GetEdgeWeight(from, to string) float64
	SetEdgeWeight(from, to string, weight float64) bool

	ShortestPath(from, to string, f ...FilterFunc) []*Vertex
	GoDFS(from string, f func(v Vertex) interface{})
	GoBFS(from string, f func(v Vertex) interface{})
}

// New returns a graph
func New() Interface {
	return (&graph{}).
		SetDirected(false).
		SetDefaultWeight(1).
		Clear()
}

func (g *graph) SetDirected(isDirected bool) Interface {
	g.isDirected = isDirected
	return g
}

func (g *graph) SetDefaultWeight(weight float64) Interface {
	g.defaultWeight = weight
	return g
}

func (g *graph) Clear() Interface {
	g.data = adjacencyList{}
	return g
}

func (g *graph) Add(id string, v interface{}) bool {
	if g.Has(id) || id == "" {
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

func (g graph) Get(id string) *Vertex {
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

func (g *graph) AddEdge(from, to string, weight ...float64) bool {
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

func (g graph) GetEdge(from, to string) *Edge {
	v := g.Get(from)
	if v == nil {
		return nil
	}
	return v.getEdge(to)
}

func (g *graph) SetEdgeWeight(from, to string, weight float64) bool {
	if e := g.GetEdge(from, to); e != nil {
		e.weight = weight
		return true
	}
	return false
}

func (g graph) GetEdgeWeight(from, to string) float64 {
	if e := g.GetEdge(from, to); e != nil {
		return e.weight
	}
	return 0
}

func (g graph) ShortestPath(from, to string, f ...FilterFunc) []*Vertex {
	hasFilter := len(f) > 0
	isExcluded := func(v *Vertex) bool { return hasFilter && !f[0](*v) }
	seen := pqueue.New()
	dist := map[string]float64{}
	prev := map[string]string{}
	path := []*Vertex{}

	if g.Get(from) == nil || g.Get(to) == nil {
		return path
	}

	for id := range g.data {
		if isExcluded(g.Get(id)) {
			fmt.Println(id)
			continue
		}

		if id == from {
			dist[from] = 0
			seen.Enqueue(from, 0)
		} else {
			dist[id] = math.Inf(1)
			seen.Enqueue(id, math.Inf(1))
		}

		prev[id] = ""
	}

	for curr := seen.Dequeue(); curr != nil; curr = seen.Dequeue() {
		smallest := curr.(string)

		if smallest == to {
			// fill path array
			for prev[smallest] != "" {
				path = append(path, g.Get(smallest))
				smallest = prev[smallest]
			}
			path = append(path, g.Get(from))
			reverseInPlace(path)

			break
		}

		g.Get(smallest).walkEdges(func(e Edge) {
			if isExcluded(g.Get(e.to)) {
				return
			}

			if newDist := dist[smallest] + e.weight; newDist < dist[e.to] {
				dist[e.to] = newDist
				prev[e.to] = smallest
				seen.Enqueue(e.to, newDist)
			}
		})
	}

	return path
}

func (g graph) GoBFS(from string, f func(Vertex) interface{}) {
	if g.Get(from) == nil {
		return
	}

	seen := queue.New()
	done := map[string]bool{}

	seen.Enqueue(from)
	done[from] = true

	for curr := seen.Dequeue(); curr != nil; curr = seen.Dequeue() {
		id := curr.(string)
		vx := g.Get(id)

		vx.walkEdges(func(e Edge) {
			if _, exists := done[e.to]; !exists {
				seen.Enqueue(e.to)
				done[e.to] = true
			}
		})

		f(*vx)
	}
}

func (g graph) GoDFS(from string, f func(Vertex) interface{}) {
	if g.Get(from) == nil {
		return
	}
	done := map[string]bool{}
	dfsRecurse(g, from, f, done)
}
