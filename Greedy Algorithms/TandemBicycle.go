package main

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// O(nlog(n)) timr | O(1) space
func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	// Write your code here.
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	totalSpeed := 0

	if fastest == true {
		for i := 0; i < len(redShirtSpeeds); i++ {
			redShirtSpeed := redShirtSpeeds[i]
			blueShirtSpeed := blueShirtSpeeds[len(blueShirtSpeeds) - (i + 1)]

			totalSpeed = totalSpeed + max(redShirtSpeed, blueShirtSpeed)
		}
	} else {		
		for i := 0; i < len(redShirtSpeeds); i++ {
			redShirtSpeed := redShirtSpeeds[i]
			blueShirtSpeed := blueShirtSpeeds[i]

			totalSpeed = totalSpeed + max(redShirtSpeed, blueShirtSpeed)
		}
	}

	return totalSpeed
}
