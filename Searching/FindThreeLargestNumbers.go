package main

import "math"

// O(n) time | O(1)	space
func FindThreeLargestNumbers(array []int) []int {
	// Write your code here.
	threeLargestNumbers := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for _, number := range array {
		updateLargestNumbers(threeLargestNumbers, number)
	}

	return threeLargestNumbers
}

func updateLargestNumbers(threeLargestNumbers []int, number int) {
	if number > threeLargestNumbers[2] {
		sortThruple(threeLargestNumbers, number, 2)
	} else if number > threeLargestNumbers[1] {
		sortThruple(threeLargestNumbers, number, 1)
	} else if number > threeLargestNumbers[0] {
		sortThruple(threeLargestNumbers, number, 0)
	}
}

func sortThruple(threeLargestNumbers []int, number, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			threeLargestNumbers[i] = number
		} else {
			threeLargestNumbers[i] = threeLargestNumbers[i+1]
		}
	}
}