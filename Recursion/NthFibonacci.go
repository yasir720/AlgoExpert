package main

import "fmt"

// O(n) time | O(1) space
func GetNthFib(n int) int {
	// Write your code here.
	lastTwo := []int{0, 1}

	if n ==1 {
		return 0
	}

	for i := 2; i <= n-1; i++ {
		nextFib := lastTwo[0] + lastTwo[1]
		lastTwo[0] = lastTwo[1]
		lastTwo[1] = nextFib
	}
	return lastTwo[1]
}

func main() {
	result := GetNthFib(6)
	fmt.Println(result)
}