package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	currentNode := linkedList
	for currentNode != nil {
		nextNewNode := currentNode.Next
		for nextNewNode != nil && nextNewNode.Value == currentNode.Value {
			nextNewNode = nextNewNode.Next
		}
		currentNode.Next = nextNewNode // this line "fixes" the linked list
		currentNode = nextNewNode // this line moves us over to the next node
														  // in the linked list
	}
		
	return linkedList
}
