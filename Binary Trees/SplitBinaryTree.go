package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | O(h) space - where n is the number of nodes in the tree and h is the height of the tree
func SplitBinaryTree(tree *BinaryTree) int {
	// Write your code here.
    treeSum := getTreeSum(tree)
	if treeSum % 2 != 0 {
		return 0
	}

	desiredSubtreeSum := treeSum / 2
	_, canBeSplit := trySubtrees(tree, desiredSubtreeSum)
	if canBeSplit {
		return desiredSubtreeSum
	}
	return 0
}

func trySubtrees(tree *BinaryTree, desiredSubtreeSum int) (int, bool) {
	if tree == nil {
		return 0, false
	}

	leftSum, leftCanBeSplit := trySubtrees(tree.Left, desiredSubtreeSum)
	rightSum, rightCanBeSplit := trySubtrees(tree.Right, desiredSubtreeSum)

	currentTreeSum := tree.Value + leftSum + rightSum
	canBeSplit := leftCanBeSplit || rightCanBeSplit || currentTreeSum == desiredSubtreeSum

	return currentTreeSum, canBeSplit
}

func getTreeSum(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}
	
	return tree.Value + getTreeSum(tree.Left) + getTreeSum(tree.Right)
}