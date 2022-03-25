package main

import "sort"

// O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array)
	leftIdx, rightIdx := 0, len(array)-1

	for leftIdx < rightIdx {
		currentSum := array[leftIdx] + array[rightIdx]

		if currentSum == target {
			return []int{array[leftIdx], array[rightIdx]}
		} else if currentSum < target {
			leftIdx = leftIdx + 1
		} else {
			rightIdx = rightIdx - 1
		}
	}
	return []int{}
}
