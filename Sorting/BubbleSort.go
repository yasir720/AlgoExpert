package insertionsort

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	sortedSlice := BubbleSort(slice)
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
func BubbleSort(array []int) []int {
	// Write your code here.
	size := len(array)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if array[i] < array[j] {
				var temp int
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	return array
}
