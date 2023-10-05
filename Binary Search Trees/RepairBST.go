package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// O(n) time | O(h) space - where n is the number of nodes in the tree and h is the height of the tree
func RepairBst(tree *BST) *BST {
	// Write your code here.
	var nodeOne *BST
	var nodeTwo *BST
	var PrevNode *BST

	var inOrderTraversal func(node *BST)
	inOrderTraversal = func(node *BST) {
		if node == nil {
			return
		}

		inOrderTraversal(node.Left)
		if PrevNode != nil && PrevNode.Value > node.Value {
			if nodeOne == nil {
				nodeOne = PrevNode
			}
			nodeTwo = node
		}

		PrevNode = node
		inOrderTraversal(node.Right)
	}

	inOrderTraversal(tree)

	nodeOne.Value, nodeTwo.Value = nodeTwo.Value, nodeOne.Value
	return tree
}
