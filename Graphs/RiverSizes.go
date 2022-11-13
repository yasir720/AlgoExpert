package main

// O(wh) time | O(wh) space
func RiverSizes(matrix [][]int) []int {
	// Write your code here.
	sizes := []int{} // we create the array to hold the sizes of the rivers found

	visited := make([][]bool, len(matrix)) // use an aux 2D array boolean (same size as input matrix) to keep track of visted nodes

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for row := 0; row < len(matrix); row++ { // iterate through the rows of the matrix
		for col := 0; col < len(matrix[row]); col++ { // iterate through the coulumbs of the matrix

			if visited[row][col] { // skip seaching any visted node as it can been search already
				continue
			}

			sizes = traverseNode(row, col, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(row, col int, matrix [][]int, visited [][]bool, sizes []int) []int {
	currentRiverSize := 0
	nodesToExplore := [][]int{{row, col}} // create a list of nodes that we'll seach as part of the inout 
										  // node given and add any neighbor that we'd also like to search

	for len(nodesToExplore) > 0 { 
		currentNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		row, col := currentNode[0], currentNode[1]

		if visited[row][col] { // skip any visited node we come accross
			continue
		}

		visited[row][col] = true // mark current node as visited
		if matrix[row][col] == 0 { // stop seaching node if it's land
			continue
		}

		currentRiverSize += 1 // increment current river size 
		unvisitedNeighbors := getNeighbors(matrix, row, col, visited) // Get the neighbors of the current node

		for _, neighbor := range unvisitedNeighbors { // add current node's niegbors to list of nodes to search
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	if currentRiverSize > 0 { // add length of current river to list
		sizes = append(sizes, currentRiverSize)
	}

	return sizes
}

func getNeighbors(matrix [][]int, row, col int, visited [][]bool) [][]int { // get the possible neighbors of a node

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
