package queue

import (
	"testing"

	tx "github.com/gregoryalbouy/go-datastructures/testx"
)

func TestQueue(t *testing.T) {
	q := newTestQueue()
	tc := tx.Testcase{
		Desc:     "queue operations",
		Input:    SliceOf(q),
		Expected: []interface{}{"Zero", "One"},
	}
	tx.Check(t, tc, tc.Input)
}

func newTestQueue() Queue {
	q := New()
	q.Dequeue()
	q.Enqueue("null")
	q.Enqueue("Zero")
	q.Dequeue()
	q.Enqueue("One")
	return q
}
