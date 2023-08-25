package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func FindLoop(head *LinkedList) *LinkedList {
	// Write your code here.
	first := head.Next
	second := first.Next

	for first != second {
		first = first.Next // moves 1 node per iteration
		second = second.Next.Next // moves 2 node per iteration
	}
	first = head

	for first != second {
		first = first.Next
		second = second.Next
	}
	return first
}

//EXPLANATION
// Let D be the distance between the start of the linked list and the
// origin of the loop in the list. Let P be distance between the origin
// of the loop and the node N where the first and second pointers overlap
// (going in the primary direction of the list). By the time the pointers
// reach N, the first pointer will have traveled a distance of length D + P.
// and the second pointer will have traveled a distance of length 2D + 2P, since
// it will have traveled twice as much as the first pointer. Thus, the distance
// between N and the origin of the loop (going in the primary direction of the
// list) can be arithmetically deduced to be 2D + 2P - D - 2P = D. With both
// pointers D length away from the origin of the loop.