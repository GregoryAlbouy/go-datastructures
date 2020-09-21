package binarysearchtree

import (
	"encoding/json"

	"github.com/gregoryalbouy/go-datastructures/queue"
)

// Tree interface
type Tree interface {
	Insert(v float64) bool
	Has(v float64) bool
	JSON() (string, error)

	root() *node
	traverseInOrder(acc *[]float64, curr *node)
	traversePreOrder(acc *[]float64, curr *node)
	traversePostOrder(acc *[]float64, curr *node)
}

type node struct {
	Value float64 `json:"value"`
	Left  *node   `json:"left,omitempty"`
	Right *node   `json:"right,omitempty"`
}

func newNode(v float64) *node {
	return &node{v, nil, nil}
}

type binarySearchTree struct {
	Root *node `json:"root"`
}

// New returns a binary search tree.
func New() Tree {
	return &binarySearchTree{}
}

func (t *binarySearchTree) Insert(v float64) bool {
	node := newNode(v)

	if t.Root == nil {
		t.Root = node
		return true
	}

	for curr := t.Root; curr != nil; {
		diff := node.Value - curr.Value
		if diff < 0 {
			if curr.Left == nil {
				curr.Left = node
				return true
			}
			curr = curr.Left
		} else if diff > 0 {
			if curr.Right == nil {
				curr.Right = node
				return true
			}
			curr = curr.Right
		} else {
			break
		}
	}
	return false
}

func (t *binarySearchTree) Has(v float64) bool {
	for curr := t.Root; curr != nil; {
		diff := v - curr.Value
		if diff == 0 {
			return true
		}
		if diff < 0 {
			curr = curr.Left
		}
		if diff > 0 {
			curr = curr.Right
		}
	}
	return false
}

func (t *binarySearchTree) JSON() (string, error) {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (t *binarySearchTree) root() *node {
	return t.Root
}

/*
SliceOfBFS returns a slice of Tree values using Breadth First Search method.
*/
func SliceOfBFS(t Tree) []float64 {
	slice := []float64{}

	if t.root() == nil {
		return slice
	}

	curr := t.root()
	q := queue.New()
	q.Enqueue(curr)

	for q.First() != nil {
		curr := q.Dequeue().(*node)
		v := curr.Value
		slice = append(slice, v)
		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}
		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}
	}
	return slice
}
