package main

import "fmt"

// O(n) time | O(1) space - space is 1 since the length of characterCont can be max 26
func FirstNonRepeatingCharacter(str string) int {
	// Write your code here.
	characterCount := map[rune]int{}

	for _, character := range str {
		characterCount[character] ++
	}

	for idx, character := range str {
		if characterCount[character] == 1{
			return idx
		}
	}
	return -1
}

func main() {
	testCharacters := ""
	result := FirstNonRepeatingCharacter(testCharacters)
	fmt.Println(result)
}