package main

func NextGreaterElement(array []int) []int {
	// Write your code here.
	result := []int{}

	for range array {
		result = append(result, -1)
	}

	stack := []int{}

	for idx := 0; idx < 2*len(array); idx++ {
		circularIdx := idx % len(array)

		for len(stack) > 0 && array[stack[len(stack)-1]] < array[circularIdx] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = array[circularIdx]
		}

		stack = append(stack, circularIdx)
	}

	return result
}
