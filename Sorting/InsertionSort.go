package insertionsort

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sortedSlice := InsertionSort(slice)
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

// best: O(n) time | O(1) space
// average: O(n^2) time | O(1) space
// worst: O(n^2) time | O(1) space
func InsertionSort(array []int) []int {
	var n = len(array)
	for i := 1; i < n; i++ {
		var j = i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
	return array
}
