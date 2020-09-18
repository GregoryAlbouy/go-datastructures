package doublylinkedlist

import (
	"testing"

	tx "github.com/gregoryalbouy/go-datastructures/testx"
)

func TestDLLReduce(t *testing.T) {
	tc := tx.Testcase{Desc: "reduce func", Expected: 6}
	l := (&doublyLinkedList{}).Push(1).Push(2).Push(3)
	reduceFunc := func(acc, curr interface{}) interface{} { return acc.(int) * curr.(int) }
	reduceInit := 1
	got := l.Reduce(reduceFunc, reduceInit)

	tx.Check(t, tc, got)
}

func TestDLLMap(t *testing.T) {
	tc := tx.Testcase{Desc: "map func", Expected: []interface{}{2, 4, 6}}
	l := (&doublyLinkedList{}).Push(1).Push(2).Push(3)
	mapFunc := func(v interface{}) interface{} { return 2 * v.(int) }
	got := SliceOf(l.Map(mapFunc))

	tx.Check(t, tc, got)
}

func TestDLLFilter(t *testing.T) {
	tc := tx.Testcase{Desc: "filter func", Expected: []interface{}{"One"}}
	l := newTestDLL()
	filterFunc := func(v interface{}) bool { return v == "One" }
	got := SliceOf(l.Filter(filterFunc))

	tx.Check(t, tc, got)
}

func TestDLLOps(t *testing.T) {
	l := newTestDLL()
	testcases := []tx.Testcase{
		{
			Desc:     "Push(\"TEST\")",
			Input:    func() []interface{} { l.Push("TEST"); return SliceOf(l) },
			Expected: []interface{}{"Zero", "One", "Two", "Three", "TEST"},
		}, {
			Desc:     "Pop()",
			Input:    func() []interface{} { l.Pop(); return SliceOf(l) },
			Expected: []interface{}{"Zero", "One", "Two", "Three"},
		}, {
			Desc:     "Unshift(\"TEST\")",
			Input:    func() []interface{} { l.Unshift("TEST"); return SliceOf(l) },
			Expected: []interface{}{"TEST", "Zero", "One", "Two", "Three"},
		}, {
			Desc:     "Shift()",
			Input:    func() []interface{} { l.Shift(); return SliceOf(l) },
			Expected: []interface{}{"Zero", "One", "Two", "Three"},
		},
	}

	for _, tc := range testcases {
		got := tc.Input.(func() []interface{})()
		tx.Check(t, tc, got)
	}
}

func newTestDLL() List {
	return New().Push("Zero").Push("One").Push("Two").Push("Three")
}

// func double(n int) int {
// 	if n == 2 {
// 		return -1
// 	}
// 	return 2 * n
// }

// func TestDummy(t *testing.T) {
// 	testcases := []testx.Testcase{
// 		{Desc: "seventeen", Input: 17, Expected: 34},
// 		{Desc: "twelve", Input: 12, Expected: 24},
// 		{Desc: "two plus two", Input: 2, Expected: 4},
// 	}
// 	tx := testx.New(testcases)

// 	for _, tc := range testcases {
// 		in := double(tc.Input.(int))
// 		tx.Check(t, tc, in)
// 	}
// }
