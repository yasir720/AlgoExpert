package main

import (
	"fmt"
	"strconv"
)

func isValidSection (section []string) bool {
	// take in a sectino of the ip address and vaerify it
	// follows the rules

	is4Digits := len(section) > 0 && len(section) < 5
	sectionNumber := strconv.Atoi(section[])
	isInRange := sectionNumber >= 0 && sectionNumber <= 255
	return is4Digits && isInRange
}

func validDotPlacements (str string) []int {
	validPlacements := make([]int, 0)

	return validPlacements
}

func main() {
	fmt.Println("hey there")
}