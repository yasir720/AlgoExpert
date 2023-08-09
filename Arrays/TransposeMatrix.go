package main

// O(wh) time | O(wh) space
func TransposeMatrix(matrix [][]int) [][]int {
	// Write your code here.
	transposeMatrix := make([][]int, len(matrix[0]))

	for col :=0; col < len(matrix[0]); col++ {
		newRow := make([]int, len(matrix))

		for row := 0; row < len(matrix); row++ {
			newRow[row] = matrix[row][col]
		}
		transposeMatrix[col] = newRow
	}
	return transposeMatrix
}
