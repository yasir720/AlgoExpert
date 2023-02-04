package main

import "math"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	// Write your code here.
	_, maxSum := findMaxSum(tree)
	return maxSum
}

func findMaxSum(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, math.MinInt32
	}

	leftMaxSumAsBranch, leftMaxPathSum := findMaxSum(tree.Left)
	rightMaxSumAsBranch, rightMaxPathSum := findMaxSum(tree.Right)
	maxChildSumAsBranch := max(leftMaxSumAsBranch, rightMaxSumAsBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumAsBranch+value+rightMaxSumAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)

	return maxSumAsBranch, maxPathSum
}

func max (nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
	}