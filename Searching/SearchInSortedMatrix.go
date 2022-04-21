package main

// O(m + n) time | O(1) space
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	// Write your code here.
	row, col := 0, len(matrix[0])-1 // start at the top right
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col = col - 1
		} else if matrix[row][col] < target {
			row = row + 1
		} else {
			return []int{row, col}
		}
	}

	return []int{-1, -1}
}
