package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n + m) time | O(1) space - where n is the length of the first Linked List, and m is the length of the second Linked List
func MergingLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// Write your code here.
	curr1, curr2 := linkedListOne, linkedListTwo

	for curr1 != curr2 {
		if curr1 == nil {
			curr1 = linkedListTwo
		} else {
			curr1 = curr1.Next
		}

		if curr2 == nil {
			curr2 = linkedListOne
		} else {
			curr2 = curr2.Next
		}
	}
	return curr1
}
