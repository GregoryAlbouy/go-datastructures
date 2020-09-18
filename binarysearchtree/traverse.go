package binarysearchtree

type traverseMethod int
type traverseFunc func(*[]float64, *node)
type traverseMap map[traverseMethod]traverseFunc

const (
	// InOrder traverse method retrieves values of a binary search tree
	// in ascending order
	InOrder traverseMethod = iota

	// PreOrder traverse method retrieves values of a binary search tree
	PreOrder

	// PostOrder traverse methods retrieves values of a binary search tree
	PostOrder
)

// SliceOfDFS returns a slice of tree values. The order depends on the method
// used: bst.InOrder, bst.PreOrder, or bst.PostOrder.
func SliceOfDFS(t Tree, method traverseMethod) []float64 {
	slice := []float64{}
	traverse := traverseMap{
		InOrder:   t.traverseInOrder,
		PreOrder:  t.traversePreOrder,
		PostOrder: t.traversePostOrder,
	}
	traverse[method](&slice, t.root())
	return slice
}

func (t *binarySearchTree) traverseInOrder(acc *[]float64, curr *node) {
	if curr.Left != nil {
		t.traverseInOrder(acc, curr.Left)
	}
	*acc = append(*acc, curr.Value)
	if curr.Right != nil {
		t.traverseInOrder(acc, curr.Right)
	}
}

func (t *binarySearchTree) traversePreOrder(acc *[]float64, curr *node) {
	*acc = append(*acc, curr.Value)
	if curr.Left != nil {
		t.traversePreOrder(acc, curr.Left)
	}
	if curr.Right != nil {
		t.traversePreOrder(acc, curr.Right)
	}
}

func (t *binarySearchTree) traversePostOrder(acc *[]float64, curr *node) {
	if curr.Left != nil {
		t.traversePostOrder(acc, curr.Left)
	}
	if curr.Right != nil {
		t.traversePostOrder(acc, curr.Right)
	}
	*acc = append(*acc, curr.Value)
}
