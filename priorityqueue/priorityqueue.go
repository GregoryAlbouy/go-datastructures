package priorityqueue

import "github.com/gregoryalbouy/go-datastructures/binaryheap"

type priorityQueue struct {
	data binaryheap.Interface
}

// Interface for priority queue
type Interface interface {
	Len() int
	Enqueue(value interface{}, priority int) Interface
	Dequeue() interface{}
	Peek() interface{}
}

type node struct {
	value    interface{}
	priority int
}

// New returns a Priority Queue
func New() Interface {
	minCompareFunc := func(A, B interface{}) int {
		a := A.(node)
		b := B.(node)

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

func (q priorityQueue) Len() int {
	return q.data.Len()
}

func (q *priorityQueue) Enqueue(v interface{}, p int) Interface {
	q.data.Insert(node{v, p})
	return q
}

func (q *priorityQueue) Dequeue() interface{} {
	node := q.data.Shift().(node)
	return node.value
}

func (q priorityQueue) Peek() interface{} {
	node := q.data.Peek().(node)
	return node.value
}
