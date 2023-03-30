package main

import "math"

// O(n^2) time | O(n) space - where n is the length of the input array
func MaxSumIncreasingSubsequence(array []int) (int, []int) {
	// Write your code here.
	sequences := make([]int, len(array))
	sums := make([]int, len(array))

	for i := range sequences {
		sequences[i] = math.MinInt32
		sums[i] = array[i]
	}

	maxSumIdx := 0

	for i, currentNum := range array {
		for j := 0; j< i; j++ {
			otherNum := array[j]
			if otherNum < currentNum && sums[j] + currentNum >= sums[i] {
				sums[i] = sums[j] + currentNum
				sequences[i] = j
			}
		}

		if sums[i] > sums[maxSumIdx] {
			maxSumIdx = i
		}
	}

	maxSum := sums[maxSumIdx]
	sequence := buildSequence(array, sequences, maxSumIdx)
	return maxSum, sequence
}

func buildSequence(array []int, sequences []int, idx int) []int {
	sequence := []int{}
	for idx != math.MinInt32 {
		sequence = append(sequence, array[idx])
		idx = sequences[idx]
	}
	reverse(sequence)
	return sequence
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}
