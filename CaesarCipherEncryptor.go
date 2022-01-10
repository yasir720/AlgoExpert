package main

import (
	"fmt"
	"strings"
)

// O(n) time | O(n) space
func CaesarCipherEncryptor(str string, key int) string {
	// Write your code here.
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for idx, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return "" // char not in alphabet
		}
		newLetter := (index + key) % 26
		runes[idx] = rune(alphabet[newLetter])
	}
	return string(runes)
}

func main() {
	test := "xyz"
	key := 2
	result := CaesarCipherEncryptor(test, key)

	fmt.Println(result)
}