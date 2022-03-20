package main

// O(wh) time | O(wh) space
func RemoveIslands(matrix [][]int) [][]int {
	// Write your code here.

	// we create a new matrix which will let us know where boarder islands are.
	// use of a 2D slice is made where all entries are zero-valued to false
	onesConnectedToBoarder := make([][]bool, len(matrix))
	for i := range matrix {
		onesConnectedToBoarder[i] = make([]bool, len(matrix[0]))
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			// define all members that are apart of the boarder
			rowIsBoarder := row == 0 || row == len(matrix)-1
			colIsBoarder := col == 0 || col == len(matrix[row])-1
			isBoader := rowIsBoarder || colIsBoarder

			if !isBoader { // skip non-boarder members
				continue
			}

			if matrix[row][col] != 1 { // skip members that are not possible islands
				continue
			}

			findOnesConnectedToBoarder(
				matrix, row, col, onesConnectedToBoarder)
		}
	}

	for row := 0; row < len(matrix)-1; row++ {
		for col := 0; col < len(matrix[row])-1; col++ {
			if onesConnectedToBoarder[row][col] {
				continue
			}

			matrix[row][col] = 0
		}
	}
	return matrix
}

// this recurrsive function checks to see if any boardering 1's are connected to
// any other 1's to form islands
func findOnesConnectedToBoarder(
	matrix[][]int, startRow, startCol int, onesConnectedToBoarder[][]bool) {
		stack := [][]int{{startRow, startCol}}
		var currentPosition []int

		for len(stack) > 0 {
			// position changed to top of stack and stack is popped
			currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]
			currentRow, currentCol := currentPosition[0], currentPosition[1]

			alreadyVisited := onesConnectedToBoarder[currentRow][currentCol]

			if alreadyVisited {
				continue
			}

			onesConnectedToBoarder[currentRow][currentCol] = true

			neighbors := getNeighbors(matrix, currentRow, currentCol)

			for _, neighbor := range neighbors {
				row, col := neighbor[0], neighbor[1]

				if matrix[row][col] != 1 {
					continue
				}

				stack = append(stack, neighbor)
			}
		}
}

// returns a 2D matrix with the location of the neighbors to a given member
// in a matrix.
func getNeighbors(matrix [][]int, row, col int)	[][]int {
	neighbors := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[row])

	// we seach the members of the matrix around the given member
	if row-1 >= 0 { // above
		neighbors = append(neighbors, []int{row - 1, col})
	}

	if row+1 < numRows { // below
		neighbors = append(neighbors, []int{row + 1, col})
	}

	if col-1 >= 0 { // to the left
		neighbors = append(neighbors, []int{col - 1, col})
	}

	if col+1 < numCols { // to the right
		neighbors = append(neighbors, []int{col + 1, col})
	}
	return neighbors
}