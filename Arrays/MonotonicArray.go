package main

import (
	"fmt"
	"math/rand"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1)
	}
	return slice
}

// O(n) time | O(1) space
func IsMonotonic(array []int) bool {
	// Write your code here.
	if len(array) < 2 {
		return true
	}

	isNonIncreasing := true
	isNonDecreasing := true

	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonIncreasing = false
		}
		if array[i] > array[i-1] {
			isNonDecreasing = false
		}
	}
	
	return isNonIncreasing || isNonDecreasing
}

func main () {
	slice := generateSlice(10)
	fmt.Println(slice)
	test := IsMonotonic(slice)
	fmt.Println(test)
}