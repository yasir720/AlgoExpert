package main

// O(wh) time | O(wh) space
func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	sizes := []int{}
	
	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}



	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {

			if visited[row][col] {
				continue
			}
			
			sizes = traverseNode(row, col, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {
	currentRiverSize := 0
	nodesToExplore := [][]int{i, j}

	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j := currentNode[0], currentNode[1]

		if visited[i][j] {
			continue
		}

		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}

		currentRiverSize += 1
		unvisitedNeighbors := findOnesConnectedToBoarder(matrix, i, j, visited)

		for _, neighbor := range unvisitedNeighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}

	return sizes
}


func getNeighbors(matrix [][]int, row, col int, visited [][]bool)	[][]int {


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