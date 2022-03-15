package main

import (
	"fmt"
	"math/rand"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(5)
	}	
	return slice
}

func absVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	sortedSquares := make([]int, len(array))

	smallerValIdx := 0
	biggerValIdx := len(array) - 1

	for i := len(array) - 1; i >= 0; i-- {
		smallerVal := array[smallerValIdx]
		biggerVal := array[biggerValIdx]

		if absVal(smallerVal) > absVal(biggerVal) {
			sortedSquares[i] = smallerVal * smallerVal
			smallerValIdx++
		} else {
			sortedSquares[i] = biggerVal * biggerVal
			biggerValIdx--
		}		
	}
	return sortedSquares
}

func main () {
	slice := generateSlice(5)
	fmt.Println(slice)
	test := SortedSquaredArray(slice)
	fmt.Println(test)
}