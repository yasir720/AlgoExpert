package main

// O(n) time | O(n) space
func ArrayOfProducts(array []int) []int {
	// Write your code here.
	products := make([]int, len(array))
	runningProduct := 1

	for idx, value := range array {
		runningProduct = runningProduct * value
		products[idx] = 1
	}

	for idx, value := range array {
		if value == 0 {
			products[idx] = 0
		} else {
			products[idx] = runningProduct / value
		}
	}
	return products
}

// import to note that this implementation fails
// the following test case:
// "array": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
