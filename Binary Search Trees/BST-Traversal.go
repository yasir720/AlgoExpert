package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// left node -> current node -> right node.
// from min to max.
// O(n) time | O(n) space
func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.

	if tree.Left != nil { // go all the way to the left
		array = tree.Left.InOrderTraverse(array)
	}

	array = append(array, tree.Value) // we're only appending here

	if tree.Right != nil { // now we check on the right side 
		array = tree.Right.InOrderTraverse(array)
	}
	return array
}

// current node -> left node -> right node.
// O(n) time | O(n) space
func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.

	array = append(array, tree.Value)
	if tree.Left != nil {
		array = tree.Left.PreOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PreOrderTraverse(array)
	}
	return array
}

// left node -> right node -> current node.
// O(n) time | O(n) space
func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	
	if tree.Left != nil {
		array = tree.Left.PostOrderTraverse(array)
	}
	if tree.Right != nil {
		array = tree.Right.PostOrderTraverse(array)
	}
	array = append(array, tree.Value)
	return array
}
