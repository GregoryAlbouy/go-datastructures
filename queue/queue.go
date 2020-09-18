package queue

// Queue type
type Queue interface {
	Len() int
	First() *node
	Enqueue(v interface{}) int
	Dequeue() interface{}
}

type node struct {
	value interface{}
	next  *node
}

func newNode(v interface{}) *node {
	return &node{v, nil}
}

type queue struct {
	length int
	first  *node
	last   *node
}

// New returns a Queue
func New() Queue {
	return &queue{}
}

// SliceOf returns a slice of values in a queue.
func SliceOf(queue Queue) []interface{} {
	slice := []interface{}{}

	for curr := queue.First(); curr != nil; curr = curr.next {
		slice = append(slice, curr.value)
	}

	return slice
}

func (q *queue) Len() int {
	return q.length
}

func (q *queue) First() *node {
	return q.first
}

func (q *queue) Enqueue(v interface{}) int {
	node := newNode(v)

	if q.length == 0 {
		q.first = node
		q.last = node
	} else {
		q.last.next = node
		q.last = node
	}

	q.length++
	return q.length
}

func (q *queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	v := q.first.value

	if q.length == 1 {
		q.first = nil
		q.last = nil
	} else {
		q.first = q.first.next
	}

	q.length--
	return v
}
