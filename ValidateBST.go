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
	return tree.validateBst(math.MinInt32, math.MaxInt32) // this set it to the max
                                                         // possible int32 value.
}

func (tree *BST) validateBst(minValue, maxValue int) bool {
    if tree == nil {
        return true
    }

    if tree.Value < minValue || tree.Value >= maxValue {
        return false // meaning that the value of the node is out of the allowed range
    }

    // we now check the possible child nodes to check validity
    leftIsValid := tree.Left.validateBst(minValue, tree.Value)
    rightIsValid := tree.Right.validateBst(tree.Value, maxValue)

    return leftIsValid && rightIsValid
}
