package main

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	// Write your code here.
	ways := make([]int, n+1)

	ways[0] = 1

	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount ++ {
			if denom <= amount {
				ways[amount] += ways[amount - denom]
			}
		}
	}
	return ways[n]
}
