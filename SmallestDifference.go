package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// Generates a slice of size, size filled with random numbers
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(5)
	}	
	return slice
}

// O(nlog(n) + mlog(m)) time | O(1) space
func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	idx1, idx2 := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for idx1 < len(array1) && idx2 < len(array2) {
		first, second := array1[idx1], array2[idx2]
		if first < second {
			current = second - first
			idx1 += 1
		} else if second < first {
			current = first - second
			idx2 += 1
		} else {
			return []int{first, second}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}

func main () {
	slice1 := generateSlice(10)
	slice2 := generateSlice(10)
	fmt.Println(slice1)
	fmt.Println(slice2)
	test := SmallestDifference(slice1, slice2)
	fmt.Println(test)
}