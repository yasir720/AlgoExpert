package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n + m) time | O(1) space
func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	// Write your code here.
	p1 := headOne
	var p1Prev *LinkedList = nil 
	p2 := headTwo

	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			p1Prev = p1
			p1 = p1.Next
		} else {
			if p1Prev != nil { // we do this as p1Prev is nil at first and can't do p1Prev.Next
				p1Prev.Next = p2
			}
			p1Prev = p2
			p2 = p2.Next
			p1Prev.Next = p1
		}
	}

	if p1 == nil {
		p1Prev.Next = p2
	}

	if headOne.Value < headTwo.Value {
		return headOne
	}
	return headTwo
}