package main

import "fmt"

// O(n) time | O(1) space
func ValidStartingCity(distances []int, fuel []int, mpg int) int {
	// Write your code here.
	numberOfCities := len(distances)
	milesRemaining := 0

	idxOfStartingCityCanidate := 0
	milesRemainingAtStartingCityCanidate := 0

	for cityIdx := 1; cityIdx < numberOfCities; cityIdx++ {
		distanceFromPreviousCity := distances[cityIdx-1]
		fuelFromPreviousCity := fuel[cityIdx-1]
		milesRemaining = milesRemaining + (fuelFromPreviousCity*mpg - distanceFromPreviousCity)

		if milesRemaining < milesRemainingAtStartingCityCanidate {
			milesRemainingAtStartingCityCanidate = milesRemaining
			idxOfStartingCityCanidate = cityIdx
		}
	}

	return idxOfStartingCityCanidate
}

func mian() {
	testDistances := []int{5, 25, 15, 10, 15}
	testFuel := []int{1, 2, 1, 0, 3}
	testMpg := 10

	result := ValidStartingCity(testDistances, testFuel, testMpg)

	fmt.Println(result)
}