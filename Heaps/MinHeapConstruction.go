package main

// Do not edit the class below except for the buildHeap,
// siftDown, siftUp, peek, remove, and insert methods.
// Feel free to add new properties and methods to the class.
type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	// Do not edit the lines below.
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

// O(n) time | O(1) space
func (h *MinHeap) BuildHeap(array []int) {
	// Write your code here.
	first := (len(array) - 2) / 2
	for currentIdx := first+1; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, len(array)-1)
	}
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftDown(currentIndex, endIndex int) {
	// Write your code here.
	childOneIdx := currentIndex*2 + 1
	for childOneIdx <= endIndex {
		childTwoIdx := -1

		if currentIndex*2 + 2 <= endIndex {
			childTwoIdx = currentIndex*2 + 2
		}

		indexToSwap := childOneIdx

		if childTwoIdx > -1 && (*h)[childTwoIdx] < (*h)[childOneIdx] {
			indexToSwap = childTwoIdx
		}

		if (*h)[indexToSwap] < (*h)[currentIndex] {
			h.swap(currentIndex, indexToSwap)
			currentIndex = indexToSwap
			childOneIdx = currentIndex*2 + 1
		} else {
			return
		}
	}
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftUp() {
	// Write your code here.
	currentIdx := h.legnth() - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 {
		current, parent := (*h)[currentIdx], (*h)[parentIdx]

		if current < parent {
			h.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
			parentIdx = (currentIdx - 1) / 2
		} else {
			return
		}
	}
}
// O(1) time | O(1) space
func (h MinHeap) Peek() int {
	// Write your code here.
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

// O(log(n)) time | O(1) space
func (h *MinHeap) Remove() int {
	// Write your code here.
	l := h.legnth()
	h.swap(0, l-1)
	peeked := (*h)[l-1]
	*h = (*h)[0 : l-1]
	h.siftDown(0, l-2)
	return peeked
}

// O(log(n)) time | O(1) space
func (h *MinHeap) Insert(value int) {
	// Write your code here.
	*h = append(*h, value)
	h.siftUp()
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) legnth() int {
	return len(h)
}
