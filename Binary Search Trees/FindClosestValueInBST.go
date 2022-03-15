package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func AbsDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func (tree *BST) FindClosestValueHelper(target, closest int) int {
	currentNode := tree

	for currentNode != nil {
		if AbsDiff(target, closest) > AbsDiff(target, currentNode.Value) {
			closest = currentNode.Value
		}
		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}
	return closest
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return tree.FindClosestValueHelper(target, tree.Value)
}

func main() {
	test := AbsDiff(10, 15)
	fmt.Println(test)
}