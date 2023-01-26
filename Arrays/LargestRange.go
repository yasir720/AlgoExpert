package main

// O(n) time | O(n) space
func LargestRange(array []int) []int {
	// Write your code here.
	best := []int{}
	longestLength := 0
	nums := map[int]bool{}

	for _, num := range array {
		nums[num] = true
	}

	for _, num := range array {
		if !nums[num] {
			continue
		}
		nums[num] = false
		currentLength, left, right := 1, num-1, num+1

		for nums[left] {
			nums[left] = false
			currentLength += 1
			left -= 1
		}
		for nums[right] {
			nums[right] = false
			currentLength += 1
			right += 1
		}

		if currentLength > longestLength {
			longestLength = currentLength
			best = []int{left+1, right-1}
		}
	}
	return best
}
