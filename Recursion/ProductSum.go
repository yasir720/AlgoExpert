package main

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array SpecialArray) int {
	// Write your code here.
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0

	for _, element := range array {
		if cast, ok := element.(SpecialArray); ok {
			sum = sum + helper(cast, multiplier + 1)
		} else if cast, ok := element.(int); ok {
			sum = sum + cast
		}
	}
	return sum * multiplier
}

func main() {
	t := "hey there"
	cast, ok := t.(SpecialArray)

}