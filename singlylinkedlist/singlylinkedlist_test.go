package main

import (
	"fmt"
	"testing"

	tx "github.com/gregoryalbouy/go-datastructures/testx"
)

func TestSSLReverse(t *testing.T) {
	fmt.Println(testList().Reverse().ToSlice())
}

func TestSLLValueAt(t *testing.T) {
	l := testList()
	testcases := []tx.Testcase{
		{Desc: "middle", Input: 2, Expected: "two"},
		{Desc: "head", Input: 0, Expected: "zero"},
		{Desc: "tail", Input: 3, Expected: "three"},
	}

	for _, tc := range testcases {
		in := tc.Input.(int)
		got := l.ValueAt(in)
		tx.Check(t, tc, got)
	}
}

func TestSLLOps(t *testing.T) {
	list := NewSinglyLinkedList()

	list.Push("null").Push("null").Pop().Push("Un").Shift().Push("Trois").Unshift("null")
	list.Set(0, "Zéro")
	list.Insert(2, "Deux")
	list.Insert(1, "null")
	list.Remove(1)
	list.Remove(3)
	list.Insert(3, "Trois")

	tc := tx.Testcase{
		Desc:     "",
		Input:    list.ToSlice(),
		Expected: []interface{}{"Zéro", "Un", "Deux", "Trois"},
	}
	tx.Check(t, tc, tc.Input)
}

func printList(l SinglyLinkedList) {
	fmt.Printf("length: %v\nhead data: %v\ntail data: %v\n", l.length, l.head.value, l.tail.value)
	curr := l.head
	for curr != nil {
		fmt.Print(curr.value, " ")
		curr = curr.next
	}
	fmt.Println()
}

func testList() *SinglyLinkedList {
	return NewSinglyLinkedList("zero", "one", "two", "three")
}
