package main

// O(n) time | O(n) space - where n is the length of the input array
func ArrayOfProducts(array []int) []int {
	// Write your code here.
	products := make([]int, len(array))
	runningProduct := 1

	// We first find the total sum of all the elements
	for idx, value := range array {
		runningProduct = runningProduct * value
		products[idx] = 1 // initalize all elements of the result array to 1
	}

	// We then calculate the respective product for each element
	for idx, value := range array {
		if value == 0 {
			products[idx] = 0
		} else {
			products[idx] = runningProduct / value
		}
	}
	return products
}