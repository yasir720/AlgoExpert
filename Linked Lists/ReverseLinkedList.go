package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func ReverseLinkedList(head *LinkedList) *LinkedList {
	// Write your code here.
	var prevNode, currentNode *LinkedList = nil, head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	return prevNode
}