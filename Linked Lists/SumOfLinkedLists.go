package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(max(n, m)) time | O(max(n, m))
func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	newLinkedList := &LinkedList{Value: 0} // this acts as a dummy node which points to the
										   // head of the answer linked list
	currentNode := newLinkedList
	carry := 0

	nodeOne := linkedListOne
	nodeTwo := linkedListTwo

	for nodeOne != nil || nodeTwo != nil || carry != 0 {
		var valueOne int
		var valueTwo int

		if nodeOne != nil {
			valueOne = nodeOne.Value
		}

		if nodeTwo != nil {
			valueTwo = nodeTwo.Value
		}

		sumOfValues := valueOne + valueTwo + carry

		newValue := sumOfValues % 10
		newNode := & LinkedList{Value: newValue}

		currentNode.Next = newNode
		currentNode = newNode

		carry = sumOfValues / 10

		if nodeOne != nil {
			nodeOne = nodeOne.Next
		}

		if nodeTwo != nil {
			nodeTwo = nodeTwo.Next
		}
	}
	return newLinkedList.Next
}
