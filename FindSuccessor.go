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
	if node.Right != nil { // the first case where the target node has a right sub-tree
		return getLeftMostChild(node.Right)
	}

	// the case where there is no right sub-tree and the following must be true
		// the target node would have to be part of the successor's left sub-tree
		// the succssor cannot be a node who's right child is the target node
		// the successor will then be the right most parent or nil if there isn't one
	return getRightMostParent(node)
}

func getLeftMostChild(node *BinaryTree) *BinaryTree {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}

func getRightMostParent(node *BinaryTree) *BinaryTree {
	currentNode := node
	// currentNode has a parent of who's right child is currentNode
	for currentNode.Parent != nil && currentNode.Parent.Right == currentNode {
		currentNode = currentNode.Parent
	}
	// when the currentnode is the left child of the parent node
	return currentNode.Parent
}