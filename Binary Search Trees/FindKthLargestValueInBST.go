package main

// This is an input class. Do not edit.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type treeInfo struct {
	numberOgNodesVisited int
	latestNodeVisited int
}

// O(h + k) time | O(h) space - h is teh height and k is the Kth value given
func FindKthLargestValueInBst(tree *BST, k int) int {
	// Write your code here.
	treeInfo := treeInfo{0, -1}
	reverseInOrderTraverse(tree, k, &treeInfo)

	return treeInfo.latestNodeVisited
}

func reverseInOrderTraverse(node *BST, k int, treeInfo *treeInfo) {
	if node == nil ||treeInfo.numberOgNodesVisited >= k {
		return
	}

	reverseInOrderTraverse(node.Right, k, treeInfo)

	if treeInfo.numberOgNodesVisited < k {
		treeInfo.numberOgNodesVisited += 1
		treeInfo.latestNodeVisited = node.Value
		reverseInOrderTraverse(node.Left, k, treeInfo)
	}
}