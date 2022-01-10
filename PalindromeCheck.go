package main

import "fmt"

// O(n) time | O(1) space
func IsPalindrome(str string) bool {
	// Write your code here.
	if len(str) == 1 {
		return true
	}

	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	testString1 := "abcdcba"
	testString2 := "asdffdsa"
	testString3 := "asdfsa"
	testString4 := "a"

	fmt.Println(IsPalindrome(testString1))
	fmt.Println(IsPalindrome(testString2))
	fmt.Println(IsPalindrome(testString3))
	fmt.Println(IsPalindrome(testString4))
}