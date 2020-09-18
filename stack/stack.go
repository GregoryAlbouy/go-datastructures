package stack

// Stack type
type Stack interface {
	Len() int
	First() interface{}

	Push(v interface{}) Stack
	Pop() interface{}
}

type node struct {
	value interface{}
	next  *node
}

func newNode(v interface{}, next *node) *node {
	return &node{v, next}
}

type stack struct {
	length int
	first  *node
	last   *node
}

// New returns a Stack
func New() Stack {
	return &stack{}
}

func (s stack) Len() int {
	return s.length
}

func (s stack) First() interface{} {
	return s.first.value
}

func (s *stack) Push(v interface{}) Stack {
	node := newNode(v, s.first)
	if s.length == 0 {
		s.last = node
	}
	s.first = node
	s.length++
	return s
}

func (s *stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	v := s.first

	if s.length == 1 {
		s.first = nil
		s.last = nil
	} else {
		s.first = s.first.next
	}

	s.length--
	return v
}
