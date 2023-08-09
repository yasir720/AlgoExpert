package main

// O(n) time | O(1) space - where n is the number of elements in the array
func MajorityElement(array []int) int {
	// Write your code here.
	count := 0
	var answer int

	for _, value := range array {
		if count == 0 {
			answer = value
		}

		if value == answer {
			count++
		} else {
			count--
		}
	}
	return answer
}
