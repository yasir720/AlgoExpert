package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sortedSlice := SelectionSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", sortedSlice)
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// best: O(n^2) time | O(1) space
// average: O(n^2) time | O(1) space
// worst: O(n^2) time | O(1) space
func SelectionSort(array []int) []int {
	// Write your code here.
	var n = len(array)
	for curtrentIdx := 0; curtrentIdx < n; curtrentIdx++ {
		var minIdx = curtrentIdx
		for j := curtrentIdx; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[curtrentIdx], array[minIdx] = array[minIdx], array[curtrentIdx]
	}
	return array
}
