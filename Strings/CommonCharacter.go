package main

// O(n * m) time | O(m) space - where n is the number of strings, and m is the length of the longest string
func CommonCharacters(strings []string) []string {
	// Write your code here.
	characterCounts := map[rune]int{}

	for _, str := range strings {
		uniqueStringCharacters := map[rune]bool{}

		for _, char := range str {
			uniqueStringCharacters[char] = true
		}

		for char := range uniqueStringCharacters {
			characterCounts[char] += 1
		}
	}

	finalCharacters := make([]string, 0)

	for char, count := range characterCounts {
		if count == len(strings) {
			finalCharacters = append(finalCharacters, string(char))
		}
	}
	return finalCharacters
}
