package main

import "sort"

// O(n*log(n)) time | O(n) space - where n is number of dishes
func SweetAndSavory(dishes []int, target int) []int {
	// Write your code here.
	sweetDishes := make([]int, 0)
	savoryDishes := make([]int, 0)

	for _, dish := range dishes {
		if dish < 0 {
			sweetDishes = append(sweetDishes, dish)
		} else if dish >0 {
			savoryDishes = append(savoryDishes, dish)
		}
	}

	sort.Slice(sweetDishes, func (i, j int) bool  {
		return sweetDishes[i] > sweetDishes[j]
		
	})

	sort.Slice(savoryDishes, func (i, j int) bool  {
		return savoryDishes[i] < savoryDishes[j]
		
	})

	bestPair := []int{0, 0}
	bestDifference := int(^uint(0) >> 1) // initalizes to max int value
	sweetIdx := 0
	savoryIdx := 0

	for sweetIdx < len(sweetDishes) && savoryIdx < len(savoryDishes) {
		currentSum := sweetDishes[sweetIdx] + savoryDishes[savoryIdx]

		if currentSum <= target {
			currentDifference := target - currentSum
			if currentDifference < bestDifference {
				bestDifference = currentDifference
				bestPair[0] = sweetDishes[sweetIdx]
				bestPair[1] = savoryDishes[savoryIdx]
			}
			savoryIdx++
		} else {
			sweetIdx++
		}
	}
	return bestPair
}
