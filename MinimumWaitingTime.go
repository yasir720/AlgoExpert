package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(9)
	}	
	return slice
}

// O(nlogn) time | O(1) space
func MinimumWaitingTime(queries []int) int {
	// Write your code here.
	totalWaitTime := 0

	if len(queries) == 1 {
		return totalWaitTime
	}

	sort.Ints(queries)

	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		totalWaitTime = totalWaitTime + (duration * queriesLeft)
	}
	return totalWaitTime
}

func main() {
	slice1 := generateSlice(6)
	fmt.Println(slice1)

	sort.Ints(slice1)
	fmt.Println(slice1)

	first := slice1[len(slice1)-1:]
	fmt.Println(first)
	rest := slice1[:len(slice1)-1]
	fmt.Println(rest)
	full := append(first, rest...)

	fmt.Println(full)

	// sort.Ints(slice1)
	// fmt.Println(slice1)

	// slice1[0], slice1[len(slice1)-1] = slice1[len(slice1)-1], slice1[0]
	// fmt.Println(slice1)


}