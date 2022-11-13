package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	// Write your code here.
	depthOne := getDescendantDepth(descendantOne, top)
	depthTwo := getDescendantDepth(descendantTwo, top)

	if depthOne > depthTwo { // case that descendantOne is lower on the tree
		depthDifference := depthOne - depthTwo
		return backtrackAncestralTree(descendantOne, descendantTwo, depthDifference)
	}

	// case that descendantTwo is lower on the tree or both descendants are at the same depth
	depthDifference := depthTwo - depthOne
	return backtrackAncestralTree(descendantTwo, descendantOne, depthDifference)
}

// iterate through the tree untill we reach the top. Keep track of the deth as we go
func getDescendantDepth(descendant, topAncestor *AncestralTree) int {
	depth := 0
	for descendant != topAncestor {
		depth = depth + 1
		descendant = descendant.Ancestor
	}

	return depth
}

func backtrackAncestralTree(lowerDescendant, higherDescendant *AncestralTree, depthDifference int) *AncestralTree {
	// bring lowerDescendant up to the same depth as higherDescendant
	for depthDifference > 0 {
		lowerDescendant = lowerDescendant.Ancestor
		depthDifference = depthDifference - 1
	}

	// keep going up the tree untill we find the first comman ancestor.
	// this works since both will be the same once the ancestor is found
	for lowerDescendant != higherDescendant {
		lowerDescendant = lowerDescendant.Ancestor
		higherDescendant = higherDescendant.Ancestor
	}

	return lowerDescendant
}
