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

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)

	var currentChangeCreated int = 0
	for _, coin := range coins {
		if coin > currentChangeCreated + 1{
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
		
	}
	return currentChangeCreated + 1
}

func main() {
	slice := generateSlice(5)
	fmt.Println(slice)
	result := NonConstructibleChange(slice)
	fmt.Println(result)

}