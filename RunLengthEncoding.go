package main

import (
	"fmt"
	"strconv"
)

// O(n) time | O(n) space
func RunLengthEncoding(str string) string {
	// Write your code here.
	previousChar := str[0]
	charRunLength := 1
	encodedStringCharacters := []byte{}

	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		if currentChar != previousChar || charRunLength == 9 {
			encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(charRunLength)[0])
			encodedStringCharacters = append(encodedStringCharacters, previousChar)

			charRunLength = 0
			previousChar = currentChar
		}
		charRunLength += 1
	}

	encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(charRunLength)[0])
	encodedStringCharacters = append(encodedStringCharacters, previousChar)

	return string(encodedStringCharacters)
}


func main() {
	string := "ABCDEFA"
	result := RunLengthEncoding(string)
	fmt.Println(result)
}