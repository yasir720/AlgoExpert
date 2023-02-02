package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Average: O(log(n)) time | O(1) space - where n is the number of nodes in the BST
// Worst: O(n) time | O(1) space - where n is the number of nodes in the BST
func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return tree.FindClosestValueHelper(target, tree.Value)
}

func (tree *BST) FindClosestValueHelper(target, closest int) int {
	currentNode := tree // we first start at the root
	closestValue := tree.Value

	for currentNode != nil {

		if AbsDiff(target, closestValue) > AbsDiff(target, currentNode.Value) {
			closestValue = currentNode.Value
		}

		// traverse the nodes of the tree to get the closest value to the target
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return closestValue
}

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
