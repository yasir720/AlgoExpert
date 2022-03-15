package main

func LongestPeak(array []int) int {
	// Write your code here.
	longestPeakLength := 0
	i := 1

	for i < len(array) - 1 {
		// to find the root of a peak
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			i = i + 1
			continue
		}

		// find the length of the left side of the peak
		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx + 1] {
			leftIdx = leftIdx - 1
		}

		// find the length of the right side of the peak
		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx - 1] {
			rightIdx = rightIdx + 1
		}

		currentLength := rightIdx - leftIdx - 1
		if currentLength > longestPeakLength {
			longestPeakLength = currentLength
		}

		// start searching for more peaks after the end of this one. Saves some time
		i = rightIdx
	}
	return longestPeakLength
}
