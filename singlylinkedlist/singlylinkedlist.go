package main

// singlyLinkedListNode represents a node in a singly linked list.
type singlyLinkedListNode struct {
	value interface{}
	next  *singlyLinkedListNode
}

// newSinglyLinkedListNode returns a *singlyLinkedListNode initialized.
// with given value and next.
func newSinglyLinkedListNode(v interface{}, next *singlyLinkedListNode) *singlyLinkedListNode {
	return &singlyLinkedListNode{v, next}
}

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	length int
	head   *singlyLinkedListNode
	tail   *singlyLinkedListNode
}

// NewSinglyLinkedList returns a singly linked list.
func NewSinglyLinkedList(values ...interface{}) *SinglyLinkedList {
	l := &SinglyLinkedList{}

	for _, v := range values {
		l.Push(v)
	}
	return l
}

// Push adds an element at the end of the list and returns the list.
func (l *SinglyLinkedList) Push(v interface{}) *SinglyLinkedList {
	node := newSinglyLinkedListNode(v, nil)

	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}

	l.length++

	return l
}

// Pop removes the last element of the list and returns the list.
func (l *SinglyLinkedList) Pop() *SinglyLinkedList {
	if l.length == 0 {
		return l
	}

	curr := l.head
	newTail := curr
	for curr.next != nil {
		newTail = curr
		curr = curr.next
	}

	newTail.next = nil
	l.tail = newTail

	l.length--

	if l.length == 0 {
		l.head = nil
		l.tail = nil
	}

	return l
}

// Shift removes the first element of the list and returns the list.
func (l *SinglyLinkedList) Shift() *SinglyLinkedList {
	if l.length == 0 {
		return l
	}

	l.head = l.head.next
	l.length--

	return l
}

// Unshift adds an element at the beginning of the list and returns the list.
func (l *SinglyLinkedList) Unshift(v interface{}) *SinglyLinkedList {
	if l.length == 0 {
		return l.Push(v)
	}

	node := newSinglyLinkedListNode(v, l.head)
	l.head = node
	l.length++

	return l
}

// get retrieves a node from its position in the list.
func (l *SinglyLinkedList) get(pos int) *singlyLinkedListNode {
	if pos < 0 || pos >= l.length {
		return nil
	}

	curr := l.head
	for i := 0; i < pos; i++ {
		curr = curr.next
	}

	return curr
}

// ValueAt returns the value of the node at given position.
func (l *SinglyLinkedList) ValueAt(pos int) interface{} {
	node := l.get(pos)

	if node == nil {
		return nil
	}

	return node.value
}

// Set sets the value of the node at given position to the given value
// and returns true. It returns false if the given position is incorrect.
func (l *SinglyLinkedList) Set(pos int, v interface{}) bool {
	node := l.get(pos)

	if node == nil {
		return false
	}

	node.value = v
	return true
}

// Insert inserts a node with given value at given position and returns
// true. It returns false if the given position is incorrect.
func (l *SinglyLinkedList) Insert(pos int, v interface{}) bool {
	prev := l.get(pos - 1)

	if prev == nil {
		return false
	}

	node := newSinglyLinkedListNode(v, prev.next)
	prev.next = node
	l.length++

	return true
}

// Remove removes a node at given position and returns true. It returns
// false if given position is incorrect.
func (l *SinglyLinkedList) Remove(pos int) bool {
	switch {
	case pos < 0 || pos >= l.length:
		return false
	case pos == 0:
		l.Shift()
		return true
	case pos == l.length-1:
		l.Pop()
		return true
	}

	prev := l.get(pos - 1)
	target := prev.next
	prev.next = target.next
	l.length--

	return true
}

// Reverse reverses the singly linked list.
func (l *SinglyLinkedList) Reverse() *SinglyLinkedList {
	var curr, prev, next *singlyLinkedListNode
	curr = l.head
	l.head = l.tail
	l.tail = l.head

	for i := 0; i < l.length; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return l
}

// ToSlice returns a slice of the values in the list. The order is conserved.
func (l *SinglyLinkedList) ToSlice() []interface{} {
	var slice []interface{}

	if l.length == 0 {
		return slice
	}

	curr := l.head
	for curr != nil {
		slice = append(slice, curr.value)
		curr = curr.next
	}

	return slice
}
