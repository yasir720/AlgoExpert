package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hey there")
	var i int = 35
	var j int = 5

	fmt.Println(i * j)
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum > target {
			right--
		} else {
			left++
		}		
	}
	return []int{}
}