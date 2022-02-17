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

func MinimumWaitingTime(queries []int) int {
	// Write your code here.
	waitTime := 0
	if len(queries) == 1 {
		return waitTime
	}

	sort.Ints(queries)

	first := queries[len(queries)-1:]
	rest := queries[:len(queries)-1]
	full := append(first, rest...)

	for i := 1; i < len(full) -1; i++ {
		waitTime += full[i]
	}
	return waitTime
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