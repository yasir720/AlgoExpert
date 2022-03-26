package main

import "sort"

func GroupAnagrams(words []string) [][]string {
	// Write your code here.
	anagrams := map[string][]string{}

	for _, word := range words {
		sortedWord := sortWord(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	result := [][]string{}

	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)
	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})
	return string(wordBytes)
}