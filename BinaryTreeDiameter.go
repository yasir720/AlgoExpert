package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	diameter int
	height int
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	// Write your code here.
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeInfo.height + rightTreeInfo.height
	maxDiameterSoFar := max(leftTreeInfo.diameter, rightTreeInfo.diameter)
	currentDiameter := max(longestPathThroughRoot, maxDiameterSoFar)
	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)

	return TreeInfo{currentDiameter, currentHeight}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
