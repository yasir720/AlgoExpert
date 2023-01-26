package main

func SunsetViews(buildings []int, direction string) []int {
	// Write your code here.
	buildingsWithSunsetViews := make([]int, 0)

	// init our start direction and step direction accordingly
	startIdx := len(buildings) - 1
	incrementDirection := -1
	if direction == "WEST" {
		startIdx = 0
		incrementDirection = 1
	}

	idx := startIdx
	runningMaxHeight := 0
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		if buildingHeight > runningMaxHeight {
			buildingsWithSunsetViews = append(buildingsWithSunsetViews, idx)
		}

		runningMaxHeight = max(runningMaxHeight, buildingHeight)

		idx += incrementDirection
	}

	if direction == "EAST" {
		reverse(buildingsWithSunsetViews)
	}

	return buildingsWithSunsetViews
}

// returns the max of two input integers
func max(a, b int) int {
	if b > a {
		return b
	}

	return a
}

// reverses the input integer array
func reverse(array []int) {
	left, right := 0, len(array)-1

	for left < right {
		array[left], array[right] = array[right], array[left]
		left += 1
		right -= 1
	}

}
