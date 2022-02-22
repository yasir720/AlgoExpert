package main

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	// Write your code here.
	if tree == nil {
		return
	}
	tree.Left, tree.Right = tree.Right, tree.Left
	tree.Left.InvertBinaryTree()
	tree.Right.InvertBinaryTree()
}
