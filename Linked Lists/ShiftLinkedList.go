package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space
func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	// Write your code here.
	listLength := 1
	listTail := head

	for listTail.Next != nil {
		listTail = listTail.Next
		listLength += 1
	}

	offset := abs(k) % listLength

	if offset == 0 {
		return head
	}

	newTailPosition := listLength - offset

	if k <= 0 {
		newTailPosition = offset
	}
	newTail := head

	for i := 1; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	listTail.Next = head
	return newHead
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func main() {
	fmt.Println(abs(-1))
}