package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | O(h) space -- n is the # of node in the tree.
// h is the height of the tree (logn)
func NodeDepths(root *BinaryTree) int {
	// Write your code here.	
	return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(node *BinaryTree, depth int) int {
	// Write your code here.
	if node == nil {
		return 0
	}

	return depth + NodeDepthsHelper(node.Left, depth + 1) + NodeDepthsHelper(node.Right, depth + 1)
}
