package main

// O(n) time | O(n) space
// func ArrayOfProducts(array []int) []int {
// 	// Write your code here.
// 	products := make([]int, len(array))
// 	runningProduct := 1

// 	for idx, value := range array {
// 		runningProduct = runningProduct * value
// 		products[idx] = 1
// 	}

// 	for idx, value := range array {
// 		if value == 0 {
// 			products[idx] = 0
// 		} else {
// 			products[idx] = runningProduct / value
// 		}
// 	}
// 	return products
// }

// important to note that the above implementation fails
// the following test case:
// "array": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]


// O(n) time | O(n) space
func ArrayOfProducts(array []int) []int {
	// Write your code here.
	products := make([]int, len(array))

	for idx, _ := range array {
		products[idx] = 1
	}

	leftRunningProduct := 1
	for idx, value := range array {
		products[idx] = leftRunningProduct
		leftRunningProduct = leftRunningProduct * value
	}

	rightRunningProduct := 1
	for i := len(array)-1; i >= 0; i-- {
		products[i] = products[i] * rightRunningProduct
		rightRunningProduct = rightRunningProduct * array[i]
	}
	return products
}