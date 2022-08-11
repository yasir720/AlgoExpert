package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space
func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	// Write your code here.
	first, second, counter := head, head, 1

	// we move our second pointer k elements AHEAD of the first pointer
	for counter <= k {
		second = second.Next
		counter = counter + 1
	}

	// we check if second is already at nil. This is the case where the
	// linked list had k elements in it. This means that first is pointing
	// to the element to be removed. in this case we remove the head and return
	if second == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}

	// all that is left to do now is to move the two ppinters one by one untill
	// second is at the end. we then remove first as it is pointing to the
	// node to be removed
	for second.Next != nil {
		second = second.Next
		first = first.Next
	}

	first.Next = first.Next.Next
}
