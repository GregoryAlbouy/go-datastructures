package binarysearchtree

import (
	"reflect"
	"testing"

	tx "github.com/gregoryalbouy/go-datastructures/testx"
)

func TestBSTInsert(t *testing.T) {
	tree := newTestBST()
	expected := &binarySearchTree{
		Root: &node{
			Value: 0,
			Left: &node{
				Value: -10,
				Left:  &node{Value: -15},
				Right: &node{Value: -5},
			},
			Right: &node{
				Value: 10,
				Left:  &node{Value: 5},
				Right: &node{Value: 15},
			},
		},
	}

	if !reflect.DeepEqual(expected, tree) {
		t.Error("tree insertion is incorrect")
	}
}

func TestBSTHas(t *testing.T) {
	tree := newTestBST()
	testcases := []tx.Testcase{
		{Desc: "root value", Input: 0., Expected: true},
		{Desc: "leftest value", Input: -15., Expected: true},
		{Desc: "rightest value", Input: 15., Expected: true},
		{Desc: "middle value", Input: 5., Expected: true},
		{Desc: "middle value", Input: -5., Expected: true},
		{Desc: "inexistent value", Input: 17., Expected: false},
		{Desc: "inexistent value", Input: -42., Expected: false},
	}

	for _, tc := range testcases {
		got := tree.Has(tc.Input.(float64))
		tx.Check(t, tc, got)
	}
}

func TestBSTJSON(t *testing.T) {
	expected := `{
  "root": {
    "value": 0,
    "left": {
      "value": -10,
      "left": {
        "value": -15
      },
      "right": {
        "value": -5
      }
    },
    "right": {
      "value": 10,
      "left": {
        "value": 5
      },
      "right": {
        "value": 15
      }
    }
  }
}`
	got, err := newTestBST().JSON()
	if err != nil {
		t.Error("json parsing error:", err)
	}

	if got != expected {
		t.Errorf("got unexpected json value:\n%v", got)
	}
}

func TestBSTSliceOfBSF(t *testing.T) {
	expected := []float64{0, -10, 10, -15, -5, 5, 15}
	got := SliceOfBFS(newTestBST())

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestBSTSliceOfDSF(t *testing.T) {
	tree := newTestBST()
	testcases := []tx.Testcase{
		{
			Desc:     "In Order",
			Input:    SliceOfDFS(tree, InOrder),
			Expected: []float64{-15, -10, -5, 0, 5, 10, 15},
		},
		{
			Desc:     "Pre Order",
			Input:    SliceOfDFS(tree, PreOrder),
			Expected: []float64{0, -10, -15, -5, 10, 5, 15},
		}, {
			Desc:     "Post Order",
			Input:    SliceOfDFS(tree, PostOrder),
			Expected: []float64{-15, -5, -10, 5, 15, 10, 0},
		},
	}

	for _, tc := range testcases {
		tx.Check(t, tc, tc.Input)
	}
}

/*
			  0
	    -10      10
	 -15  -5    5  15
*/
func newTestBST() Tree {
	tree := New()
	tree.Insert(0)
	tree.Insert(-10)
	tree.Insert(10)
	tree.Insert(-15)
	tree.Insert(-5)
	tree.Insert(15)
	tree.Insert(5)
	return tree
}
