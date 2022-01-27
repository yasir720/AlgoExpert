package main

import "fmt"

// O(m + n) time | O(c) space | m = length of characters & n = lenth of document &
// c = number of unique characters in character
func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	characterCount := map[rune]int{}

	for _, character := range characters {
		characterCount[character] ++
	}

	for _, character := range document {
		if characterCount[character] == 0 {
			return false
		}
		characterCount[character] --
	}
	return true
}

func main() {
	testCharacters := "$aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz--"
	fmt.Println(len(testCharacters))
	testDocument := "yasir-was-here"
	fmt.Printf("its %T\n", testDocument[0])
	result := GenerateDocument(testCharacters, testDocument)
	fmt.Println(result)
}