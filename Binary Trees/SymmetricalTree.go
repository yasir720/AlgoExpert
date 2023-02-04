package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SymmetricalTree(tree *BinaryTree) bool {
	// Write your code here.
	return checkMirrored(tree.Left, tree.Right)
}

func checkMirrored(left, right *BinaryTree) bool {
	if left != nil && right != nil && left.Value == right.Value {
		return checkMirrored(left.Left, right.Right) && checkMirrored(left.Right, right.Left)
	}

	return left == right
}


// This version here is explained but doesn't work (3/13)
// package main

// // This is an input class. Do not edit.
// type BinaryTree struct {
// 	Value int

// 	Left  *BinaryTree
// 	Right *BinaryTree
// }

// func SymmetricalTree(tree *BinaryTree) bool {
// 	// Write your code here.
// 	return checkMirrored(tree.Left, tree.Right)
// }

// func checkMirrored(left, right *BinaryTree) bool {
// 	isChildrenThere := left != nil && right != nil // check that children are in position
// 	isValuesEqual := left.Value == right.Value // check that the values are equal

// 	if  isChildrenThere && isValuesEqual {
// 		// check below level for mirrored childeren (outside and inside)
// 		isOutideChildrenMirrored := checkMirrored(left.Left, right.Right)
// 		isInideChildrenMirrored :=  checkMirrored(left.Right, right.Left)

// 		// is below level good to go?
// 		return isOutideChildrenMirrored && isInideChildrenMirrored 
// 	}

// 	return left == right // ?
// }
