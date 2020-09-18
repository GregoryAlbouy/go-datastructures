package doublylinkedlist

// List represents a doubly linked list
type List interface {
	Filter(f FilterFunc) List
	Map(f MapFunc) List
	Reduce(f ReduceFunc, init interface{}) interface{}

	Len() int
	Head() *node
	Tail() *node

	Push(v interface{}) List
	Pop() List
	Unshift(v interface{}) List
	Shift() List

	InsertAt(pos int, v interface{}) bool
	RemoveAt(pos int) bool
	ValueAt(pos int) interface{}
}

// FilterFunc takes a value parameter and returns a boolean
type FilterFunc = func(v interface{}) bool

// MapFunc type
type MapFunc = func(v interface{}) interface{}

// ReduceFunc type
type ReduceFunc = func(acc, curr interface{}) interface{}

type node struct {
	value interface{}
	prev  *node
	next  *node
}

func newNode(v interface{}, prev *node, next *node) *node {
	return &node{v, prev, next}
}

type doublyLinkedList struct {
	length int
	head   *node
	tail   *node
}

// New returns a doubly linked list
func New() List {
	return &doublyLinkedList{}
}

// Push adds an element at the end of the list
func (l *doublyLinkedList) Push(v interface{}) List {
	node := newNode(v, l.tail, nil)

	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		l.setTail(node)
	}

	l.length++

	return l

	// return l.insertBefore(l.head, v)
}

// Pop removes the last element of the list
func (l *doublyLinkedList) Pop() List {
	return l.remove(l.tail)
}

// Unshift adds an element at the beginning of the list
func (l *doublyLinkedList) Unshift(v interface{}) List {
	node := newNode(v, nil, l.head)

	if l.length == 0 {
		l.head = node
		l.tail = node
	} else {
		l.setHead(node)
	}

	l.length++

	return l
}

// Shift removes the first element of the list
func (l *doublyLinkedList) Shift() List {
	return l.remove(l.head)
}

func (l *doublyLinkedList) remove(node *node) List {
	if l.length == 0 {
		return l
	}

	if l.length == 1 {
		return l.reset()
	}

	prev := node.prev
	next := node.next

	if node == l.head {
		l.head = node.next
	} else {
		prev.next = next
	}

	if node == l.tail {
		l.tail = node.prev
	} else {
		next.prev = prev
	}

	l.length--

	return l
}

func (l *doublyLinkedList) InsertAt(pos int, v interface{}) bool {
	return true
}

func (l *doublyLinkedList) RemoveAt(pos int) bool {
	return true
}

func (l *doublyLinkedList) ValueAt(pos int) interface{} {
	return nil
}

func (l *doublyLinkedList) nodeAt(pos int) *node {
	return nil
}

func (l *doublyLinkedList) reset() List {
	l.head = nil
	l.tail = nil
	l.length = 0
	return l
}

func (l *doublyLinkedList) setHead(node *node) {
	l.head.prev = node
	l.head = node
}

func (l *doublyLinkedList) setTail(node *node) {
	l.tail.next = node
	l.tail = node
}

func (l *doublyLinkedList) Head() *node {
	return l.head
}

func (l *doublyLinkedList) Tail() *node {
	return l.tail
}

func (l *doublyLinkedList) Len() int {
	return l.length
}

func (l *doublyLinkedList) Filter(f FilterFunc) List {
	for curr := l.head; curr != nil; curr = curr.next {
		if !f(curr.value) {
			l.remove(curr)
		}
	}

	return l
}

func (l *doublyLinkedList) Map(f MapFunc) List {
	for curr := l.head; curr != nil; curr = curr.next {
		curr.value = f(curr.value)
	}

	return l
}

func (l *doublyLinkedList) Reduce(f ReduceFunc, init interface{}) interface{} {
	acc := init
	for curr := l.head; curr != nil; curr = curr.next {
		acc = f(acc, curr.value)
	}
	return acc
}

// SliceOf returns a slice of a list values.
func SliceOf(list List) []interface{} {
	slice := []interface{}{}

	if list.Len() == 0 {
		return slice
	}

	for curr := list.Head(); curr != nil; curr = curr.next {
		slice = append(slice, curr.value)
	}

	return slice
}
