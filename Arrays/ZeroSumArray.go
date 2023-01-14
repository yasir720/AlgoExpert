package main

// O(n) time | O(n) space - where n is the length of nums
func ZeroSumSubarray(nums []int) bool {
	// Write your code here.
	sums := map[int]bool{0: true} // keep track of of running sums
	runningSum := 0

	for _, num := range nums {
		runningSum += num
		// check if the running sum is included in sums
		if _, sumInSet := sums[runningSum]; sumInSet {
		return true
	}
	// if running sum is not included, add it
	sums[runningSum] = true
	}
	// no zero value subarray
	return false
}
