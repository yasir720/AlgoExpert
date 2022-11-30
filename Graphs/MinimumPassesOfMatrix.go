package main

type IntPair struct {
	First, Second int
}

// O(wh) time | O(wh) space
func MinimumPassesOfMatrix(matrix [][]int) int {
	// Write your code here.
	passes := convertNegitives(matrix)
	if !containsNegitive(matrix) { // goes over the metrix to check for any negitive positions
		return passes - 1 // we do -1 because the last iterate is a check, not a pass
	} else {
		return -1 // after we do out best to convert, if there is a negitive, then it's not able to be coinverted so we return -1
	}
}

func convertNegitives(matrix [][]int) int {
	queue := getAllPositivePositions(matrix) // will return an array of all the row,col pairs of positive positions

	var passes int = 0
	for len(queue) > 0 {
		var currentSize int = len(queue)

		for currentSize > 0 {
			// In Go, removing elements from the start of a list is an O(n)-time operation.
			// To make this an O(1)-time operation, we could use a more legitimate queue structure.
			// For our time complexity analysis, we'll assume this runs in 0(1) time.
			nextElement := queue[0] // the current position where we will check for negitive neighbors
			queue = queue[1:] // pop the element we're about to check
			currentRow, currentCol := nextElement.First, nextElement.Second

			neighborPositions := getNeighbors(matrix, currentRow, currentCol) // will return an array of neighbors [row, col]
			for _, position := range neighborPositions {
				row, col := position.First, position.Second

				value := matrix[row][col]
				if value < 0 { // if the neighbor is negitive
					matrix[row][col] *= -1 // make it positive
					queue = append(queue, IntPair{row, col}) // add this new positive position to the queue
				}
			}
			currentSize -= 1 // we do this so we know how many elements in the queue we're able to convert since newly converted positive have to wait
							 // till the next iteration to be able to conver negitives
		}
		passes += 1
	}
	return passes	
}

func getAllPositivePositions(matrix [][]int) []IntPair {
	positivePositions := make([]IntPair, 0)
	for row := range matrix {
		for col := range matrix[row] {
			value := matrix[row][col]
			if value > 0 {
				positivePositions = append(positivePositions, IntPair{row, col})
			}
		}
	}
	return positivePositions
}

func getNeighbors(matrix [][]int, row, col int) []IntPair {

	neighbors := make([]IntPair, 0)

	// we seach the members of the matrix around the given member
	if row > 0 { // above
		neighbors = append(neighbors, IntPair{row - 1, col})
	}

	if row < len(matrix)-1 { // below
		neighbors = append(neighbors, IntPair{row + 1, col})
	}

	if col  > 0 { // to the left
		neighbors = append(neighbors, IntPair{row, col - 1})
	}

	if col < len(matrix[0])-1 { // to the right
		neighbors = append(neighbors, IntPair{row, col + 1})
	}
	return neighbors
}

func containsNegitive(matrix [][]int) bool {
	for _, row := range matrix {
		for _, value := range row	{
			if value < 0 {
				return true
			}
		}
	}
	return false
}
