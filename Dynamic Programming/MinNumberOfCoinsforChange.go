package main

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Write your code here.
	numOfCoins := make([]int, n+1)
	for i := range numOfCoins {
		numOfCoins[i] = math.MaxInt32
	}

	numOfCoins[0] = 0

	for _, denom := range denoms {
		for amount := range numOfCoins {
			if denom <= amount {
				numOfCoins[amount] = min(numOfCoins[amount], numOfCoins[amount - denom]+1)
			}
		}
	}
	if numOfCoins[n] != math.MaxInt32 {
		return numOfCoins[n]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}