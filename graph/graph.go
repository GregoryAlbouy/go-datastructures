package graph

// Work in progress

import "fmt"

type vertex interface{}

type adjacencyList map[vertex][]vertex

type graph struct {
	adjacencyList adjacencyList
}

// Interface of a graph
type Interface interface {
	AddVertex(v vertex)
	AddEdge(v1, v2 vertex)
	RemoveEdge(v1, v2 vertex)
	RemoveVertex(v vertex)

	debug()
}

// New returns a graph
func New() Interface {
	return &graph{map[vertex][]vertex{}}
}

func (g *graph) AddVertex(v vertex) {
	_, exists := g.adjacencyList[v]
	if !exists {
		g.adjacencyList[v] = []vertex{}
	}
}

func (g *graph) AddEdge(v1, v2 vertex) {
	g.adjacencyList[v1] = append(g.adjacencyList[v1], v2)
	g.adjacencyList[v2] = append(g.adjacencyList[v2], v1)
}

func (g *graph) RemoveEdge(v1, v2 vertex) {
	_, v1Exists := g.adjacencyList[v1]
	_, v2Exists := g.adjacencyList[v2]

	if !v1Exists || !v2Exists {
		return
	}

	g.adjacencyList[v1] = findAndRemove(g.adjacencyList[v1], v2)
	g.adjacencyList[v2] = findAndRemove(g.adjacencyList[v2], v1)
}

func (g *graph) RemoveVertex(v vertex) {
	_, exists := g.adjacencyList[v]

	if !exists {
		return
	}

	delete(g.adjacencyList, v)

	for vertex := range g.adjacencyList {
		g.adjacencyList[vertex] = findAndRemove(g.adjacencyList[vertex], v)
	}
}

func (g graph) debug() {
	fmt.Printf("%+v\n", g.adjacencyList)
}

func findAndRemove(edgeList []vertex, v vertex) []vertex {
	newEdgeList := []vertex{}

	for _, edge := range edgeList {
		if edge != v {
			newEdgeList = append(newEdgeList, edge)
		}
	}

	return newEdgeList
}
