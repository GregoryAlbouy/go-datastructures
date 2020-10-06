package graph

import (
	"fmt"
	"testing"

	"github.com/gregoryalbouy/go-datastructures/testx"
)

func TestGraphAdd(t *testing.T) {
	// g := New().SetDirected(true)
}

func TestGraphShortestPath(t *testing.T) {
	tc := testx.Testcase{
		Desc:     "standard path",
		Input:    newTestDirectedGraph(),
		Expected: []string{"A", "B", "C", "D", "E", "F"},
	}
	path := tc.Input.(Interface).ShortestPath("A", "F")
	got := func() (ids []string) {
		for _, node := range path {
			ids = append(ids, node.id)
		}
		return
	}()

	testx.Check(t, tc, got)
}

// func TestGraphGoBFS(t *testing.T) {
// 	storeIn := func(ids []string) func(v vertex) {
// 		return func(v vertex) {
// 			ids = append(ids, v.id)
// 		}
// 	}
// 	testcases := []testx.Testcase{}
// }

func newTestDirectedGraph() Interface {
	ids := []string{"romance", "thriller", "media", "support",
		"book", "movie", "paper", "screen", "action"}
	g := New()

	for _, id := range ids {
		g.Add(id, fmt.Sprintf("%s is cool", id))
	}

	g.AddEdge("media", "book")
	g.AddEdge("media", "movie")
	g.AddEdge("media", "support")
	g.AddEdge("book", "paper")
	g.AddEdge("book", "action")
	g.AddEdge("book", "romance")
	g.AddEdge("book", "thriller")
	g.AddEdge("movie", "screen")
	g.AddEdge("movie", "action")
	g.AddEdge("movie", "romance")
	g.AddEdge("movie", "thriller")
	g.AddEdge("support", "paper")
	g.AddEdge("support", "screen")

	return g
}

/*
           A
         / |
     5 /   | 20
     / 12  |   7
   B - - - D - - - E
     \     |       |
     6 \   | 4     | 3
         \ |       |
           C - - - F
               17
*/
func newTestUndirectedGraph() Interface {
	ids := []string{"D", "B", "E", "F", "A", "C"}
	vertices := func() (vertices []struct {
		id    string
		value interface{}
	}) {
		for _, id := range ids {
			vertices = append(vertices, struct {
				id    string
				value interface{}
			}{id, "value" + id})
		}
		return
	}()
	g := New()
	g.AddMany(vertices...)
	g.AddEdge("B", "D", 12)
	g.AddEdge("F", "E", 3)
	g.AddEdge("F", "C", 17)
	g.AddEdge("A", "D", 20)
	g.AddEdge("D", "E", 7)
	g.AddEdge("D", "C", 4)
	g.AddEdge("B", "C", 6)
	g.AddEdge("A", "B", 5)
	return g
}
