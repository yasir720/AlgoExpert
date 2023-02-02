package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// O(n) time | O(d) space
func (tree *BST) ValidateBst() bool {
	// Write your code here.
	return tree.ValidateBstHelper(math.MinInt32, math.MaxInt32) // this set it to the min/max int32 value
}

func (tree *BST) ValidateBstHelper(minValue, maxValue int) bool {
    if tree == nil { // a nil node/BST is a valid node/BST
        return true
    }

    if tree.Value < minValue || tree.Value >= maxValue {
        return false // meaning that the value of the node is out of the allowed range for that node/BST
    }

    // we now check the possible child nodes to check validity
    leftIsValid := tree.Left.ValidateBstHelper(minValue, tree.Value)
    rightIsValid := tree.Right.ValidateBstHelper(tree.Value, maxValue)

    return leftIsValid && rightIsValid
}
