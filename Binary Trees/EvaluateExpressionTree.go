package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | O(h) space - where n is the number of nodes in the Binary Tree, and h is the height of the Binary Tree
func EvaluateExpressionTree(tree *BinaryTree) int {
	// Write your code here.
	if tree.Value > 0 {
		return tree.Value
	}

	leftValue := EvaluateExpressionTree(tree.Left)
	rightValue := EvaluateExpressionTree(tree.Right)

	if tree.Value == -1 {
		return leftValue + rightValue
	}
	if tree.Value == -2 {
		return leftValue - rightValue
	}
	if tree.Value == -3 {
		return leftValue / rightValue
	}
	return leftValue * rightValue // the left over case of -4
	
}
