package main

// Do not edit the class below except for the
// PopulateSuffixTrieFrom and Contains methods.
// Feel free to add new properties and methods
// to the class.
type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

// O(n^2) time | O(n^2) space
func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	// Write your code here.
	for i := range str {
		node := trie // look at type struct
		for j := i; j <len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found { // check if in map then base on conditional (found)
				node[letter] = NewSuffixTrie() // add letters that aren't already added -- new characters
			}
			node = node[letter] // move down to the next letter (just added if not already there)
		}
		node['*'] = nil // add the special end character once done with iteration
	}
}

// O(m) time | O(1) space
func (trie SuffixTrie) Contains(str string) bool {
	// Write your code here.
	node := trie // look at type struct
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, found := node[letter]; !found { // check if in map then base on conditional (found)
			return false // if the current letter was not found
		}
		node = node[letter] // move down to the next letter
	}
	_, found := node['*'] // ensure we found a suffix by checking for the special ending character
	return found
}
