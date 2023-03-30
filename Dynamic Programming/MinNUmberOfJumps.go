package main

// O(n) time | O(1) space - where n is the length of the input array
func MinNumberOfJumps(array []int) int {
	// Write your code here.
	if len(array) == 1 {
		return 0
	}

	jumps := 0
	maxReach := array[0]
	steps := array[0]

	for i:=1; i < len(array)-1; i++ {
		if i+array[i] > maxReach {
			maxReach = i + array[i]
		}
		steps -= 1
		if steps == 0 {
			jumps += 1
			steps = maxReach - i
		}
	}
	return jumps + 1
}
