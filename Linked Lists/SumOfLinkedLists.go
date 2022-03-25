package SumOfLinkedLists

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

	// carry is apart of this condition as there are cases where we still have to sum
	// values even when there is no more values in either linked list.
	// such a case is the sum of 9->9->9 & 9->9 where the result is
	// 8->9->0->1 which is longer than either given linked list
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
