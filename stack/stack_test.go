package stack

import (
	"testing"

	tx "github.com/gregoryalbouy/go-datastructures/testx"
)

func TestStackOps(t *testing.T) {
	s := newTestStack()
	testcases := []tx.Testcase{
		{
			Desc:     "Push(\"TEST\")",
			Input:    func() interface{} { s.Push("TEST"); return s.First() },
			Expected: "TEST",
		}, {
			Desc:     "Pop()",
			Input:    func() interface{} { s.Pop(); return s.First() },
			Expected: "one",
		},
	}

	for _, tc := range testcases {
		got := tc.Input.(func() interface{})()
		tx.Check(t, tc, got)
	}
}

func newTestStack() Stack {
	s := New()
	s.Push("zero")
	s.Push("one")
	return s
}
