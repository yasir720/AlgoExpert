package main

// This is the struct of the input root. Do not edit it.
// O(n) time | O(n) space -- n is the # of node in the tree
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	sums := []int{}
	BranchSumsHelper(root, 0, &sums)
	return sums
}

func BranchSumsHelper(node *BinaryTree, runningSum int, sums *[]int) {
	// Write your code here.
	if node == nil {
		return
	}

	newRunningSum := runningSum + node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, newRunningSum)
		return
	}
	// not &sums as sums here is initalized as a pointer already by the function
	BranchSumsHelper(node.Left, newRunningSum, sums)
	BranchSumsHelper(node.Right, newRunningSum, sums)
}
