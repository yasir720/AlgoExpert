package main

// O(n) time | O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	seqIdx := 0 // start at the begining of the input sequence
	for _, value := range array { // iterate through the input array
		if seqIdx == len(sequence) { // stop iterating if the sequence index is out of bounds
			break
		}
		if value == sequence[seqIdx] { // increment the sequence index if we find the current character
			seqIdx ++
		}
	}
	return seqIdx == len(sequence) // did we increment through the full sequence?
}
