package main

import (
	"fmt"
	"math/rand"
)

// Generates a slice of size, size filled with random numbers
func generateTwoDSlice(x, y int) [][]int {
	twoD := make([][]int, x)
	for i := 0; i < x; i++ {
		twoD[i] = make([]int, y)
		for j := 0; j < y; j++ {
			twoD[i][j] = rand.Intn(8)
		}		
	}
	return twoD
}

// func SpiralTraverse(array [][]int) []int {
// 	// Write your code here.
// 	outer, inner := len(array), len(array[0])
// 	numberOfElements := outer * inner
	
// 	fmt.Println(array)
// 	fmt.Println(numberOfElements)
// 	return nil
// }

// O(n) time | O(n) space - where n is the total number of elements in the array
func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
	  return []int{}
	}
	result := []int{}
	spiralFill(array, 0, len(array)-1, 0,
	  len(array[0])-1, &result)
	return result
  }
  func spiralFill(array [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	if startRow > endRow || startCol > endCol {
	  return
	}
	for col := startCol; col <= endCol; col++ {
	  *result = append(*result, array[startRow][col])
	}
	for row := startRow + 1; row <= endRow; row++ {
	  *result = append(*result, array[row][endCol])
	}
	for col := endCol - 1; col >= startCol; col-- {
	  if startRow == endRow {
		break
	  }
	  *result = append(*result, array[endRow][col])
	}
	for row := endRow - 1; row >= startRow+1; row-- {
	  if startCol == endCol {
		break
	  }
	  *result = append(*result, array[row][startCol])
	}
	
	spiralFill(array, startRow+1, endRow-1,
	  startCol+1, endCol-1, result)
  }

func main () {
	twoD := generateTwoDSlice(5,5)
	fmt.Println(twoD)
	test := SpiralTraverse(twoD)
	fmt.Println(test)
}