package main

func getNextIndex (currentIdx int, array []int) int {
	hopLength := array[currentIdx]
	nextIdx := (currentIdx + hopLength) % len(array)
	
	if nextIdx >= 0 {
		return nextIdx
	}
	return nextIdx + len(array)
}

func HasSingleCycle(array []int) bool {
	// Write your code here.
	numberOfVisitedElements := 0
	currentIdx := 0

	for numberOfVisitedElements < len(array) {
		if numberOfVisitedElements > 0 && currentIdx == 0 { // if we return back to the start of the array without
															// looping all indexies, then false
			return false
		}
		numberOfVisitedElements = numberOfVisitedElements + 1
		currentIdx = getNextIndex(currentIdx, array)
	}	
	return currentIdx == 0
}
