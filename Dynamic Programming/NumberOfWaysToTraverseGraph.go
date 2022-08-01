package main

// O(n + m) time | O(n + m) space
func NumberOfWaysToTraverseGraph(width int, height int) int {
	// Write your code here.
	NumberOfWays := make([][]int, height+1)
	
	for i := range NumberOfWays {
		NumberOfWays[i] = make([]int, width+1)
	}

	for widthIdx := 1; widthIdx < width+1; widthIdx++ {
		for heightIdx := 1; heightIdx < height+1; heightIdx++ {
			if widthIdx == 1 || heightIdx == 1 {
				NumberOfWays[heightIdx][widthIdx] = 1
			} else {
				waysLeft := NumberOfWays[heightIdx][widthIdx-1]
				waysUp := NumberOfWays[heightIdx-1][widthIdx]
				NumberOfWays[heightIdx][widthIdx] = waysLeft + waysUp
			}
		}
	}
	return NumberOfWays[height][width]
}


// another method using math trick

// O(n + m) time | O(1) space
// func NumberOfWaysToTraverseGraph(width int, height int) int {
// 	// Write your code here.
// 	xDistanceToCorner := width - 1
// 	YDistanceToCorner := height - 1

// 	numerator := factorial(xDistanceToCorner + YDistanceToCorner)
// 	denominator := factorial(xDistanceToCorner) * factorial(YDistanceToCorner)

// 	return numerator / denominator
// }

// func factorial( num int) int {
// 	result := 1

// 	for n := 2; n < num+1; n++ {
// 		result *= n
// 	}
	
// 	return result
// }