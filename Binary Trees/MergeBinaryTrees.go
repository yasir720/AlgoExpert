package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | O(h) space - where n is the # of nodes on the smaller tree and h
// is the height of the shorter tree
func MergeBinaryTrees(tree1 *BinaryTree, tree2 *BinaryTree) *BinaryTree {
	// Write your code here.
	if tree1 == nil	{
		return tree2
	}
	if tree2 == nil {
		return tree1
	}

	tree1.Value += tree2.Value
	tree1.Left = MergeBinaryTrees(tree1.Left, tree2.Left)
	tree1.Right = MergeBinaryTrees(tree1.Right, tree2.Right)
	return tree1
}
