package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	data := []int{}

	stack := []*BinarySearchTree{bst}
	visited := make(map[*BinarySearchTree]struct{})
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if node.left != nil {
			if _, ok := visited[node.left]; !ok {
				stack = append(stack, node.left)
				continue
			}
		}

		stack = stack[:len(stack)-1]
		visited[node] = struct{}{}
		data = append(data, node.data)

		if node.right != nil {
			stack = append(stack, node.right)
		}
	}

	return data
}
