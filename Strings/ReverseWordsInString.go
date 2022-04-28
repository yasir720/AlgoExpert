package main

// O(n) time | O(n) space
func ReverseWordsInString(str string) string {
	// Write your code here.
	characters := make([]byte, 0)
	for _, char := range []byte(str) {
		characters = append(characters, char)
	}
	reverseListRange(characters, 0, len(characters)-1)
	
	startOfWord := 0
	for startOfWord < len(characters) {
		endOfWord := startOfWord
		for endOfWord < len(characters) && characters[endOfWord] != ' ' {
			endOfWord = endOfWord + 1
		}
		
		reverseListRange(characters, startOfWord, endOfWord-1)
		startOfWord = endOfWord + 1
	}
	
	return string(characters)
}

func reverseListRange(list []byte, rangeStart, rangeEnd int) {
	start := rangeStart
	end := rangeEnd

	for start < end {
		temp := list[start]
		list[start] = list[end]
		list[end] = temp
		start = start + 1
		end = end - 1
	}
}