package main

import "fmt"

func main() {
	// fmt.Println("hey there")
	// var i int = 35
	// var j int = 5

	// fmt.Println(i * j)

	i := 100
	j := 200

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(&j)

}

// func TwoNumberSum(array []int, target int) []int {
// 	// Write your code here.
// 	sort.Ints(array)
// 	left, right := 0, len(array)-1
// 	for left < right {
// 		currentSum := array[left] + array[right]
// 		if currentSum == target {
// 			return []int{array[left], array[right]}
// 		} else if currentSum > target {
// 			right--
// 		} else {
// 			left++
// 		}
// 	}
// 	return []int{}
// }