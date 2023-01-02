package main

// O(nm*8^s + ws) time | O(nm + ws) space
func BoggleBoard(board [][]rune, words []string) []string {
	// Write your code here.
	return nil
}

func explore(i, j int, board [][]rune, trie Trie, visited [][]bool, finalWords map[string]bool) {

}

func getNeighbors(i, j int, board [][]rune) [][]int {
	neighbors := [][]int{}

	if i > 0 && j > 0 {
		neighbors = append(neighbors, []int{i-1, j-1})
	}
	if i > 0 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i-1, j+1})
	}
	if i < len(board)-1 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i+1, j+1})
	}
	if i < len(board)-1 && j > 0 {
		neighbors = append(neighbors, []int{i+1, j-1})
	}
	if i > 0 {
		neighbors = append(neighbors, []int{i-1, j})
	}
	if i < len(board)-1 {
		neighbors = append(neighbors, []int{i+1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, []int{i, j-1})
	}
	if j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i, j+1})
	}

	return neighbors
}

type Trie struct {
	children map[rune]Trie
	word string
}

func (t Trie) add(word string) {
	current := t

	for _, letter := range word {
		if _, found := current.children[letter]; !found {
			current.children[letter] = Trie{
				children: map[rune]Trie{},
			}
		}
		current = current.children[letter]
	}
	current.children['*'] = Trie{ // end character and we also store the word for easy access
		children: map[rune]Trie{},
		word: word,
	}
}
