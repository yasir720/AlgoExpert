package main

// O(n) time | O(1) space - where n is the number of seats
func BestSeat(seats []int) int {
	// Write your code here.
	bestSeat := -1
	maxSpace := 0

	left := 0
	for left < len(seats) {
		right := left + 1
		for right < len(seats) && seats[right] == 0 {
			right += 1
		}

		avalibleSpace := right - left -1
		if avalibleSpace > maxSpace {
			bestSeat = (left + right) / 2
			maxSpace = avalibleSpace
		}
		left = right
	}
	return bestSeat
}
