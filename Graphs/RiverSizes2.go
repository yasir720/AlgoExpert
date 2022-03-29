package main

func RiverSizes(matrix [][]int) []int {
	// Write your code here.

	// first we have to traverse the input matrix to find all possible beginings of rivers.
	// we then will ave to keep track of these visited locations as to not redo work or
	// count any rivers twice
	// when we do find any rivers, we have to check its neighbors to continuations
		// it is note worthy to make sure that these neighbors weren't prevously checked.
	//once we are done checking all possible neightbors, we can add the length of the current
	// river to the result array answer


	// this slice will serve as out answer
	lengthsOfRivers := []int{}

	// this will serve as our auxilary matrix to keep track of visited elements
	riversAlreadyVisited := make([][]bool, len(matrix))

	for i := range matrix {
		riversAlreadyVisited[i] = make([]bool, len(matrix[0]))
	}






	// 1. traverse the inout matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {

			if matrix[row][col] == 0 {
				continue
			}
			
			findOnesConnectedToBoarder(
				matrix, row, col, riversAlreadyVisited)
		}
	}
	return nil
}

// this recurrsive function checks to see if any boardering 1's are connected to
// any other 1's to form islands
func findOnesConnectedToBoarder(matrix[][]int, startRow, startCol int, onesConnectedToBoarder [][]bool) {
	stack := [][]int{{startRow, startCol}}
	var currentPosition []int

	for len(stack) > 0 {
		// current position changed to top of stack and stack is popped
		currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]

		currentRow, currentCol := currentPosition[0], currentPosition[1]

		alreadyVisited := onesConnectedToBoarder[currentRow][currentCol]

		if alreadyVisited { // skip already visited elements
			continue
		}

		// set the current element we car checking to true as it can now be considered visited
		onesConnectedToBoarder[currentRow][currentCol] = true


		// from this bit we expect to get back a 2D array holding the location
		// of all river elements that are connected to the current visited element
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
