package main

func Semordnilap(words []string) [][]string {
	// Write your code here.
	wordSet := map[string]bool{}
	for _, word := range words {
		wordSet[word] = true
	}

	SemordnilapPairs := [][]string{}
	
	for _, word := range words {
		reverse := reverse(word)
		if _, reverseInSet := wordSet[reverse]; reverseInSet && reverse != word {
			SemordnilapPairs = append(SemordnilapPairs, []string{word, reverse})
			delete(wordSet, word)
			delete(wordSet, reverse)
		}
	}

	return SemordnilapPairs
}

func reverse(s string) string {
	reversed := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	return string(reversed)
}
