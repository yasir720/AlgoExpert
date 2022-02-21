package main

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// correctly add in a node to the BST
// average: O(Log(n)) time | O(1) space
// worst: O(n) time | O(1) space
func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.

	currentNode := tree
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil { // when we reach the end of the left side then 
										// instert and stop traversing
				currentNode.Left = &BST{Value: value}
				break
			} else {// keep traversing the left side of the tree
				currentNode = currentNode.Left
			}
		} else {
			if currentNode.Right == nil { // when we reach the end of the right side then 
										 // instert and stop traversing
				currentNode.Right = &BST{Value: value}
				break
			} else { // keep traversing the right side of the tree
				currentNode = currentNode.Right
			}
		}
	}
	return tree
}

// search a BST for a given value
// average: O(Log(n)) time | O(1) space
// worst: O(n) time | O(1) space
func (tree *BST) Contains(value int) bool {
	// Write your code here.

	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
				currentNode = currentNode.Right
		} else {
			return true
		}
	}
	return false
}

// Remove a node from a BST
// average: O(Log(n)) time | O(1) space
// worst: O(n) time | O(1) space
func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.

	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(value int, parent *BST) {
	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			parent = currentNode
			currentNode = currentNode.Right
		} else {
			// case 1 where the node to be removed had 2 child nodes.
			// we must replace it with the min value of the right side of the
			// node to  be removed
			if currentNode.Left != nil && currentNode.Right != nil {
				// set hte value of the current node to the min of the right subtree
				currentNode.Value = currentNode.Right.getMinValue()
				// now remove min node of the right subtree
				currentNode.Right.remove(currentNode.Value, currentNode)
			} else if parent == nil { // case 2 where we are removing the root node
				if currentNode.Left != nil {
					currentNode.Value = currentNode.Left.Value
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				} else if currentNode.Right != nil {
					currentNode.Value = currentNode.Right.Value
					currentNode.Right = currentNode.Right.Right
					currentNode.Left = currentNode.Right.Left
				} else {
					// the case in which we are dealing with a single node tree.
					// we do nothing...
				}
			} else if parent.Left == currentNode { // case 3 where we are dealing with leaf nodes
				if currentNode.Left != nil {
					parent.Left = currentNode.Left
				} else {
					parent.Left = currentNode.Right
				}
			} else if parent.Right == currentNode {
				if currentNode.Left != nil {
					parent.Right = currentNode.Left
				} else {
					parent.Right = currentNode.Right
				}
			}
			break
		}
	}
}


func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}
