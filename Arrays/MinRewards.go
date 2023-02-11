package main

func MinRewards(scores []int) int {
	// Write your code here.
	lastMaxIdx := 0
	lastMaxReward := 1
	previousReward := 1
	totalRewards := 1

	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			totalRewards += totalRewards + previousReward + 1
			lastMaxIdx = i
			lastMaxReward = previousReward
		} else {
			if previousReward == 1 {
				totalRewards += i -lastMaxIdx
				if i-lastMaxIdx >= lastMaxReward {
					totalRewards++
				} else {
					totalRewards += 1
				}
				previousReward = 1
			}
		}
	}
	return totalRewards
}
