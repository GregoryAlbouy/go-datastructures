package graph

import "testing"

// TODO: real tests
func TestGraph(t *testing.T) {
	g := New()
	g.AddVertex("aaa")
	g.AddVertex("bbb")
	g.AddVertex("ccc")
	g.AddEdge("bbb", "ccc")
	g.RemoveEdge("bbb", "ccc")
	g.AddEdge("bbb", "ccc")
	g.AddEdge("bbb", "ccc")
	g.AddEdge("bbb", "aaa")
	g.debug()
	g.RemoveVertex("ccc")
	g.debug()
}
