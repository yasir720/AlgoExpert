package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

// O(n) time | O(d) space - where n is the total number of elements in the array,
// including sub-elements, and d is the greatest depth of "special" arrays in the array
func ProductSum(array SpecialArray) int {
	// Write your code here.
	return ProductSumHelper(array, 1)
}

func ProductSumHelper(array SpecialArray, multiplier int) int {
	sum := 0

	for _, element := range array {
		if cast, ok := element.(SpecialArray); ok {
			sum = sum + ProductSumHelper(cast, multiplier + 1)
		} else if cast, ok := element.(int); ok {
			sum = sum + cast
		}
	}
	return sum * multiplier
}