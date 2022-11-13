package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space
func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	currentNode := linkedList
	for currentNode != nil {
		nextNewNode := currentNode.Next
		for nextNewNode != nil && nextNewNode.Value == currentNode.Value {
			nextNewNode = nextNewNode.Next // move the new next node one node over
		}
		currentNode.Next = nextNewNode // if there was a duplicate, we skip that node. 
		currentNode = nextNewNode // we now move the current node into the right position
	}
		
	return linkedList
}

// easier way to go about it
// O(n) time | O(1) space
// func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
// 	// Write your code here.
// 	currentNode := linkedList
// 	for currentNode.Next != nil {
// 		if currentNode.Value == currentNode.Next.Value {
// 			currentNode.Next = currentNode.Next.Next
// 		} else {
// 			currentNode = currentNode.Next
// 		}
// 	}
		
// 	return linkedList
// }
