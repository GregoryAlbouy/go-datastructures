package binaryheap

import "github.com/gregoryalbouy/go-datastructures/compare"

var defaultCompareFunc = func(A, B interface{}) int {
	a, aIsInt := A.(int)
	b, bIsInt := B.(int)

	if !aIsInt || !bIsInt {
		panic("BinaryHeap: The default compareFunc only works on ints. In order to work with different type, it must be provided a custom compareFunc via .setCompareFunc method.")
	}

	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

// New returns a new binary heap.
func New() Interface {
	return &binaryHeap{
		compareFunc: defaultCompareFunc,
	}
}

type binaryHeap struct {
	values      []interface{}
	compareFunc compare.Func
}

// Interface for Binary Heap.
type Interface interface {
	Len() int
	Insert(v interface{}) Interface
	InsertMany(values ...interface{}) Interface
	Shift() interface{}
	Clear() Interface
	ToSlice() []interface{}
	SetCompareFunc(f compare.Func) Interface
}

func (h *binaryHeap) SetCompareFunc(f compare.Func) Interface {
	h.compareFunc = f
	return h
}

// Len returns the number of elements in the Binary Heap.
func (h binaryHeap) Len() int {
	return len(h.values)
}

// ToArray returns an array of the elements in the Binary Heap.
func (h binaryHeap) ToSlice() []interface{} {
	return h.values
}

func (h *binaryHeap) Insert(v interface{}) Interface {
	h.values = append(h.values, v)
	h.siftUp()
	return h
}

func (h *binaryHeap) InsertMany(values ...interface{}) Interface {
	for _, v := range values {
		h.Insert(v)
	}
	return h
}

func (h *binaryHeap) Shift() interface{} {
	n := h.Len()
	if n == 0 {
		return nil
	}

	first := h.values[0]
	if n == 1 {
		h.Clear()
		return first
	}

	last := h.values[n-1]
	h.values = h.values[:n-1]

	h.values[0] = last
	h.siftDown()

	return first
}

func (h *binaryHeap) Clear() Interface {
	h.values = []interface{}{}
	return h
}

func (h *binaryHeap) siftUp() {
	i := h.Len() - 1
	f := h.compareFunc

	for i > 0 {
		v := h.values[i]
		iParent := (i - 1) / 2
		vParent := h.values[iParent]

		if compare.InfOrEq(v, vParent, f) {
			break
		}

		h.swap(i, iParent)
		i = iParent
	}
}

func (h *binaryHeap) siftDown() {
	i := 0
	n := h.Len()

	for i < n {
		iLeft := 2*i + 1
		iRight := iLeft + 1
		iMax := func() int {
			switch {
			case !h.hasIndex(iLeft):
				return i
			case !h.hasIndex(iRight):
				return h.maxValueIndex(i, iLeft)
			default:
				return h.maxValueIndex(i, iLeft, iRight)
			}
		}()

		if i == iMax {
			break
		}

		h.swap(i, iMax)
		i = iMax
	}
}

func (h *binaryHeap) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h binaryHeap) hasIndex(i int) bool {
	return i < h.Len()
}

func (h binaryHeap) maxValueIndex(indexes ...int) int {
	iMax := indexes[0]

	for _, i := range indexes {
		if compare.SupOrEq(h.values[i], h.values[iMax], h.compareFunc) {
			iMax = i
		}
	}

	return iMax
}
