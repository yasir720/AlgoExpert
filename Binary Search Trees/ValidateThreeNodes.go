package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	// Write your code here.
	if isDescendant(nodeTwo, nodeOne) {
		return isDescendant(nodeThree, nodeTwo)
	}

	if isDescendant(nodeTwo, nodeThree) {
		return isDescendant(nodeOne, nodeTwo)
	}
	return false
}

func isDescendant(node, target *BST) bool {
	currentNode := node
	for currentNode != nil && currentNode != target {
		if target.Value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return currentNode == target
}