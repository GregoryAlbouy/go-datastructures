package priorityqueue

import "github.com/gregoryalbouy/go-datastructures/binaryheap"

type priorityQueue struct {
	data binaryheap.Interface
}

// Interface for priority queue
type Interface interface {
	Len() int
	Enqueue(value interface{}, priority float64) Interface
	Dequeue() interface{}
	Peek() interface{}
}

// Node is a priority queue elements It has a value and a priority.
type Node struct {
	value    interface{}
	priority float64
}

// Value returns the value of the node.
func (n Node) Value() interface{} {
	return n.value
}

// Priority returns the priority of the node.
func (n Node) Priority() float64 {
	return n.priority
}

// New returns a Priority Queue
func New() Interface {
	minCompareFunc := func(A, B interface{}) int {
		a := A.(Node)
		b := B.(Node)

		if a.priority < b.priority {
			return 1
		}
		if a.priority > b.priority {
			return -1
		}
		return 0
	}
	q := &priorityQueue{}
	q.data = binaryheap.New().SetCompareFunc(minCompareFunc)
	return q
}

// SliceOf returns a slice of a Priority Queue values.
func SliceOf(queue Interface) []interface{} {
	q := queue.(*priorityQueue)
	nodes := q.data.ToSlice()
	values := []interface{}{}

	for _, v := range nodes {
		node := v.(Node)
		values = append(values, node.value)
	}

	return values
}

func (q priorityQueue) Len() int {
	return q.data.Len()
}

func (q *priorityQueue) Enqueue(v interface{}, p float64) Interface {
	q.data.Insert(Node{v, p})
	return q
}

func (q *priorityQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}
	node := q.data.Shift().(Node)
	return node.value
}

func (q priorityQueue) Peek() interface{} {
	if q.Len() == 0 {
		return nil
	}
	node := q.data.Peek().(Node)
	return node.value
}
