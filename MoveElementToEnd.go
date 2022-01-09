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

// O(n) time | O(1) space
func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	left, right := 0, len(array)-1

	for left < right {
		for left < right && array[right] == toMove {
			right -= 1
		}
		if array[left] == toMove {
			array[left], array[right] = array[right], array[left]
		}
		left += 1
	}
	return array
}

func main () {
	slice := generateSlice(10)
	fmt.Println(slice)
	test := MoveElementToEnd(slice, 0)
	fmt.Println(test)
}