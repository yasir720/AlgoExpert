package main

import "sort"

// O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	sort.Ints(array) // sort the inout array in place <
	leftIdx, rightIdx := 0, len(array)-1 // left will be the smallest and right will be the bigest

	for leftIdx < rightIdx { // as long as the pointers don't cross
		currentSum := array[leftIdx] + array[rightIdx] // create a temp/possible sum for the current iteration

		if currentSum == target { // the case that we have reached the target sum
			return []int{array[leftIdx], array[rightIdx]}
		} else if currentSum < target { // the case where our current sum needs to increase
			leftIdx = leftIdx + 1
		} else { // the case where our current sum needs to decrease
			rightIdx = rightIdx - 1
		}
	}
	return []int{} // return empty if we don't find a possible sum in the input array
}

// faster version using a hashmap
// O(n) time | O(n) space
// func TwoNumberSum(array []int, target int) []int {
// 	// Write your code here.
// 	m := make(map[int]bool)

// 	for _, number := range array {
// 		neededNumber := target - number
// 		_, hasNeededNumber := m[neededNumber]
// 		if hasNeededNumber {
// 			return []int{neededNumber, number}
// 		} else {
// 			m[number] = true
// 		}
// 	}
// 	return []int{}
// }