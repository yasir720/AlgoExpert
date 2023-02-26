package main

func IsMonotonic(array []int) bool {
	// Write your code here.
	if len(array) < 2 {
		return true
	}

	isNonincreasing := true
	isNonDecreasing := true

	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			isNonDecreasing = false
		} else if array[i] < array[i-1] {
			isNonincreasing = false
		}
	}
	return isNonincreasing || isNonDecreasing
}
