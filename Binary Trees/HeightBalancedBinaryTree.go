package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	isBalanced bool
	height int
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	// Write your code here.
	return getTreeInfo(tree).isBalanced
}

func getTreeInfo(node *BinaryTree) TreeInfo {
	if node == nil { // base case where we have hit a leaf nodes child
		return TreeInfo{true, -1}
	}

	leftSubTreeInfo := getTreeInfo(node.Left)
	rightSubTreeInfo := getTreeInfo(node.Right)

	isBalanced := leftSubTreeInfo.isBalanced && rightSubTreeInfo.isBalanced && abs(
		leftSubTreeInfo.height - rightSubTreeInfo.height) <= 1
	height := max(leftSubTreeInfo.height, rightSubTreeInfo.height) + 1

	return TreeInfo{isBalanced, height}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
