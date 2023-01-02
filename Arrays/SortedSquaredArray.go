package main

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	sortedSquares := make([]int, len(array)) // create the return array

	smallerValIdx := 0 // smaller value or left value which is the smallest or more negitive
	biggerValIdx := len(array) - 1 // larger value or right value which is the largest or more positive

	for i := len(array) - 1; i >= 0; i-- { // increment through the input array
		smallerVal := array[smallerValIdx]
		biggerVal := array[biggerValIdx]

		if absVal(smallerVal) > absVal(biggerVal) { // find the absolute values, then square then, add to return array, and increment pointer accordingly
			sortedSquares[i] = smallerVal * smallerVal
			smallerValIdx++
		} else {
			sortedSquares[i] = biggerVal * biggerVal
			biggerValIdx--
		}		
	}
	return sortedSquares
}

func absVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
