package main

import "math"

func CollidingAsteroids(asteroids []int) []int {
	// Write your code here.
	resultStack := make([]int, 0)

	for _, asteroid := range asteroids {
		if asteroid > 0 {
			resultStack = append(resultStack, asteroid)
			continue
		}

		for {
			lastStackIdx := len(resultStack) - 1
			if len(resultStack) == 0 || resultStack[lastStackIdx] < 0 {
				resultStack = append(resultStack, asteroid)
				break
			}

			asteroidSize := int(math.Abs(float64(asteroid)))

			if resultStack[lastStackIdx] > asteroidSize {
				break
			}

			if resultStack[lastStackIdx] == asteroidSize {
				resultStack = resultStack[:lastStackIdx]
				break
			}

			resultStack = resultStack[:lastStackIdx]
		}
	}
	return resultStack
}
