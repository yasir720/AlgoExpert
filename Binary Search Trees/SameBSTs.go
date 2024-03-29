package main

// O(n^2) time | O(d) space - where n is the number of nodes in each array, 
// respectively, and d is the depth of the BST that they represent
func SameBsts(arrayOne, arrayTwo []int) bool {
	// Write your code here.
	if len(arrayOne) != len(arrayTwo) {
		return false
	}
	
	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}
	
	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	leftOne := getSmaller(arrayOne)
	leftTwo := getSmaller(arrayTwo)
	rightOne := getBiggerOrEqual(arrayOne)
	rightTwo := getBiggerOrEqual(arrayTwo)

	return SameBsts(leftOne, leftTwo) && SameBsts(rightOne, rightTwo)
}

func getSmaller(array []int) []int {
	smaller := []int{}
	for i := 1; i < len(array); i ++ {
		if array[i] < array[0] {
			smaller = append(smaller, array[i])
		}
	}
	return smaller
}

func getBiggerOrEqual(array []int) []int {
	biggerOrEqual := []int{}
	for i := 1; i < len(array); i++ {
		if array[i] >= array[0] {
			biggerOrEqual = append(biggerOrEqual, array[i])
		}
	}
	return biggerOrEqual
}