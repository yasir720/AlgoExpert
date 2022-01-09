package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(5)
	}	
	return slice
}

// O(n^2) time | O(n) space
func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	sort.Ints(array)
	triplets := [][]int{}

	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1

		for left < right {
			tripletSum := array[i] + array[left] + array[right] 

			if tripletSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if tripletSum < target {
				left += 1
			} else if tripletSum > target {
				right -= 1
			}
		}
	}
	return triplets
}

func main () {
	slice := generateSlice(10)
	fmt.Println(slice)
	test := ThreeNumberSum(slice, 5)
	fmt.Println(test)
}