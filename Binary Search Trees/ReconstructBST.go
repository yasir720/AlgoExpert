package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// O(n^2) time | O(h) space
func ReconstructBst(preOrderTraversalValues []int) *BST {
	// Write your code here.
	if len(preOrderTraversalValues) == 0 {
		return nil
	}

	currentValue := preOrderTraversalValues[0]
	rightSubTreeRootIdx := len(preOrderTraversalValues)

	for idx := 1; idx < len(preOrderTraversalValues); idx++ {
		value := preOrderTraversalValues[idx]
		if value >= currentValue {
			rightSubTreeRootIdx = idx
			break
		}
	}

	leftSubTree := ReconstructBst(preOrderTraversalValues[1:rightSubTreeRootIdx])
	rightSubTree := ReconstructBst(preOrderTraversalValues[rightSubTreeRootIdx:])

	return &BST{currentValue, leftSubTree, rightSubTree}
}
