package main

func SubarraySort(array []int) []int {
	// Write your code here.
	leftIdx := -1
	rightIdx := -1
	minValue := array[len(array)-1]
	maxValue := array[0]
	for idx, value := range array {
		if value < maxValue {
			leftIdx = idx
		} else {
			maxValue = value
		}
	}
	
	for i := len(array)-1; i >= 0; i-- {
		if array[i] > minValue {
			rightIdx = i
		} else {
			minValue = array[i]
		}
	}
	return []int{rightIdx, leftIdx}
}
