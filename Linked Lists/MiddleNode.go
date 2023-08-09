package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space - where n is the number of nodes in the linked list
func MiddleNode(linkedList *LinkedList) *LinkedList {
	// Write your code here.
	slowNode := linkedList // will reach the half point of the LL when fast reaches the end
	fastNode := linkedList

	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	return slowNode
}
