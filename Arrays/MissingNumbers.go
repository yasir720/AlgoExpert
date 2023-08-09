package main

// O(n) time | O(1) space - where n is the length of the input array
func MissingNumbers(nums []int) []int {
	// Write your code here.
	total := sum(arrayFromAToB(1, len(nums)+3)) // len(nums)+3) acts as n. the extra 1 comes from the index change since 1,n not 0,n

	for _, num := range nums {
		total -= num
	}

	averageMissingValue := total / 2
	foundFirstHalf := 0
	foundSecondHalf := 0

	for _, num := range nums {
		if num <= averageMissingValue {
			foundFirstHalf += num
		} else {
			foundSecondHalf += num
		}
	}

	expectedFirstHalf := sum(arrayFromAToB(1, averageMissingValue+1))
	expectedSecondHalf := sum(arrayFromAToB(averageMissingValue+1, len(nums)+3))
	return []int{expectedFirstHalf - foundFirstHalf, expectedSecondHalf - foundSecondHalf}
}

func arrayFromAToB(a, b int) []int {
	array := make([]int, 0)
	for num := a; num < b; num++ {
		array = append(array, num)
	}
	return array
}

func sum(array []int) int {
	total := 0
	for _, num := range array {
		total += num
	}
	return total
}

// USING BITWISE OPERATIONS
// // O(n) time | O(1) space - where n is the length of the input array
// func MissingNumbers(nums []int) []int {
// 	// Write your code here.
// 	solutionXOR := 0
// 	for i := 0; i < len(nums)+3; i++ {
// 		solutionXOR ^= i
// 		if i < len(nums) {
// 			solutionXOR ^= nums[i]
// 		}
// 	}

// 	solution := []int{0, 0}
// 	setBit := solutionXOR & -solutionXOR
// 	for i := 0; i < len(nums)+3; i++ {
// 		if (i & setBit) == 0 {
// 			solution[0] ^= i
// 		} else {
// 			solution[1] ^= i
// 		}

// 		if i < len(nums) {
// 			if (nums[i] & setBit) == 0 {
// 				solution[0] ^= nums[i]
// 			} else {
// 				solution[1] ^= nums[i]
// 			}
// 		}
// 	}

// 	if solution[0] > solution[1] {
// 		solution[0], solution[1] = solution[1], solution[0]
// 	}
// 	return solution
// }
