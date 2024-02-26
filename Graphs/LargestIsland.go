package main

type node struct {
	row, col int
}

// O (w * h) time | O(w * h) space - where w is the width of the matrix, 
// and h is the height of the matrix
func LargestIsland(matrix [][]int) int {
	// Write your code here.
	islandSizes := make([]int, 0)

	// island number starts at 2 to avoid overwriting existing 0s and 1s
	islandNumber := 2
	for row := 0; row < len(matrix); row++ {
		for col := 0; col <len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				islandSizes = append(islandSizes, getSizeFromNode(row, col, matrix, islandNumber))
				islandNumber += 1
			}
		}
	}

	maxSize := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col <len(matrix[row]); col++ {
			if matrix[row][col] != 1 {
				continue
			}

			landNeighbors := getLandNeighbors(row, col, matrix)
			islands := map[int]bool{}
			for _, neighbor := range landNeighbors {
				islands[matrix[neighbor.row][neighbor.col]] = true
			}

			size := 1
			for island := range islands {
				size += islandSizes[island - 2]
			}

			maxSize = max(maxSize, size)
		}
	}
	return maxSize
}

func getSizeFromNode(row, col int, matrix [][]int, islandNumber int) int {
	size := 0
	nodesToExplore := []node{{row, col}}
	var currNode node
	for len(nodesToExplore) > 0 {
		currNode, nodesToExplore = nodesToExplore[len(nodesToExplore)-1], nodesToExplore[:len(nodesToExplore)-1]
		currRow, currCol := currNode.row, currNode.col

		if matrix[currRow][currCol] != 0 {
			continue
		}

		matrix[currRow][currCol] = islandNumber
		size += 1
		nodesToExplore = append(nodesToExplore, getLandNeighbors(currRow, currCol, matrix)...) //https://www.geeksforgeeks.org/how-to-use-ellipsis-in-golang/
	}

	return size
}

func getLandNeighbors(row, col int, matrix [][]int) []node {
	landNeighbors := make([]node, 0)

	if row > 0 && matrix[row-1][col] != 1{
		landNeighbors = append(landNeighbors, node{row - 1, col})
	}
	if row < len(matrix)-1 && matrix[row+1][col] != 1{
		landNeighbors = append(landNeighbors, node{row + 1, col})
	}
	if col > 0 && matrix[row][col - 1] != 1{
		landNeighbors = append(landNeighbors, node{row, col - 1})
	}
	if col < len(matrix[0])-1 && matrix[row][col+1] != 1{
		landNeighbors = append(landNeighbors, node{row, col + 1})
	}

	return landNeighbors
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
