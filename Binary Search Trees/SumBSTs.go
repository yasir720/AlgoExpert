package main

import "math"

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | o(h) space - where n is the number of nodes in the
// tree and h is hte height of the tree
func SumBsts(tree *BinaryTree) int {
	// Write your code here.
	return getTreeInfo(tree).TotalSumBstNodes
}

type TreeInfo struct {
	IsBST bool
	MaxValue int
	MinValue int
	BstSum int
	BstSize int
	TotalSumBstNodes int
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{
			IsBST: true,
			MaxValue: math.MinInt64,
			MinValue: math.MaxInt64,
			BstSum: 0,
			BstSize: 0,
			TotalSumBstNodes: 0,
		}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	satisfiesBestProp := tree.Value > leftTreeInfo.MaxValue && tree.Value <= rightTreeInfo.MinValue
	isBst := satisfiesBestProp && leftTreeInfo.IsBST && rightTreeInfo.IsBST

	maxValue := int(math.Max(float64(tree.Value), math.Max(float64(leftTreeInfo.MaxValue), float64(rightTreeInfo.MaxValue))))
	minValue := int(math.Min(float64(tree.Value), math.Min(float64(leftTreeInfo.MinValue), float64(rightTreeInfo.MinValue))))

	var bstSum, bstSize int

	toalSumBstNodes := leftTreeInfo.TotalSumBstNodes + rightTreeInfo.TotalSumBstNodes

	if isBst {
		bstSum = tree.Value + leftTreeInfo.BstSum + rightTreeInfo.BstSum
		bstSize = 1 + leftTreeInfo.BstSize + rightTreeInfo.BstSize

		if bstSize >= 3 {
			toalSumBstNodes = bstSum
		}
	}

	return TreeInfo{
		IsBST: isBst,
		MaxValue: maxValue,
		MinValue: minValue,
		BstSum: bstSum,
		BstSize: bstSize,
		TotalSumBstNodes: toalSumBstNodes,
	}
}
