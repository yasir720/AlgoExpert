package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

// O(h) time | O(1) space
func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	// Write your code here.
	if node.Right != nil {
		return getLeftMostChild(node.Right)
	}

	return getRightMostParent(node)
}

func getLeftMostChild(node *BinaryTree) *BinaryTree {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}

func getLeftMgetRightMostParentostChild(node *BinaryTree) *BinaryTree {
	currentNode := node
	for currentNode.Parent != nil && currentNode.Parent.Right == currentNode {
		currentNode = currentNode.Parent
	}

	return currentNode.Parent
}