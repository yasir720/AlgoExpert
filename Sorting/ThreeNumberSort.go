package main

func ThreeNumberSort(array []int, order []int) []int {
	// Write your code here.
	firstValue := order[0]
	secondValue := order[1]
	thirdValue := order[2]

	firstIdx, secondIdx, thirdIdx := 0, 0, len(array) - 1

	for secondIdx <= thirdIdx {
		value := array[secondIdx]

		if value == firstValue {
			array[secondIdx], array[firstIdx] = array[firstIdx], array[secondIdx]
			firstIdx = firstIdx + 1
			secondIdx = secondIdx + 1
		} else if value == secondValue {
			secondIdx = secondIdx + 1
		} else if value == thirdValue {
			array[secondIdx], array[thirdIdx] = array[thirdIdx], array[secondIdx]
			thirdIdx = thirdIdx - 1
		}
	}
	
	return array
}
