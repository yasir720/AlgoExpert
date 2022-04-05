package main

import "sort"

// O(nlog(n)) time | O(1) space
func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	// Write your code here.
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	frontRowColor := "blue"
	if redShirtHeights[0] < blueShirtHeights[0] {
		frontRowColor = "red"
	}

	for i := 0; i < len(blueShirtHeights); i++ {
		if frontRowColor == "red" {
			if redShirtHeights[i] >= blueShirtHeights[i] {
				return false
			}
		} else  {
			if blueShirtHeights[i] >= redShirtHeights[i] {
				return false
			}
		}
	}
	
	return true
}
