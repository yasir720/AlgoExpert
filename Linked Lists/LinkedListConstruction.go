// Feel free to add methods and fields to the struct definitions.
package main

type Node struct { // here we create the structure for the node
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct { // we create the doubly linked list now
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	// Write your code here.
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	// Write your code here.
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
		return
	}
	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	// Write your code here.
	if ll.Tail == nil {
		ll.SetHead(node)
		return
	}
	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	// Write your code here.
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}

	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	// Write your code here.
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	ll.Remove(nodeToInsert)
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next

	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}

	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	// Write your code here.
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	node := ll.Head
	currentPosition := 1
	for node != nil && currentPosition != position {
		node = node.Next
		currentPosition = currentPosition + 1
	}
	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	// Write your code here.
	node := ll.Head
	for node != nil {
		nodeToRemove := node
		node = node.Next
		if nodeToRemove.Value == value {
			ll.Remove(nodeToRemove)
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// Write your code here.
	if node == ll.Head {
		ll.Head = ll.Head.Next
	}
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}
	ll.removeNodeLinks(node)
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	// Write your code here.
	node := ll.Head
	for node != nil && node.Value != value {
		node = node.Next
	}
	return node != nil
}
 
func (ll *DoublyLinkedList) removeNodeLinks(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}