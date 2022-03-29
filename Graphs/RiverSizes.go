package main

// O(wh) time | O(wh) space
func RiverSizes(matrix [][]int) []int {
	// Write your code here.

	// we create a new matrix which will allow us to keep track of the elements
	// that we have already visited.
	// use of a 2D slice is made where all entries are zero-valued to false
	riversAlreadyVisited := make([][]bool, len(matrix))

	for i := range matrix {
		riversAlreadyVisited[i] = make([]bool, len(matrix[0]))
	}

	// this bit is used to circle the perimiter and search for islands along the boarder
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {

			if matrix[row][col] == 0 {
				continue
			}

			//this bit will help us find the 1s connected to border 1s
			findOnesConnectedToBoarder(
				matrix, row, col, riversAlreadyVisited)
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
func findOnesConnectedToBoarder(matrix[][]int, startRow, startCol int, onesConnectedToBoarder [][]bool) {
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
		neighbors = append(neighbors, []int{row, col - 1})
	}

	if col+1 < numCols { // to the right
		neighbors = append(neighbors, []int{row, col + 1})
	}
	return neighbors
}