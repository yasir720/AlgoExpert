package main

import "fmt"

func LongestPalindromicSubstring(str string) string {
	// Write your code here.
	if len(str) == 1 {
		return str
	}

	result := ""	

	for i := 1; i < len(str); i++ {
		oddSubstring := palindromeLengthCheck(str, i-1, i+1)

		evenSubstring := palindromeLengthCheck(str, i-1, i)

		oddLength := len(oddSubstring)
		evenLength := len(evenSubstring)
		
		longest := evenSubstring

		if oddLength > evenLength {
			longest = oddSubstring
		}

		if len(longest) > len(result) {
			result = longest
		}
		//fmt.Println("i:", i, "odd sub:", oddSubstring, "even sub:", evenSubstring)
	}
	return result
}

func palindromeLengthCheck(str string, leftIdx, rightIdx int) string {
	for leftIdx > 0 && rightIdx < len(str)-1 {
		if str[leftIdx] != str[rightIdx] {
			break
		}
		leftIdx -= 1
		rightIdx += 1

		//fmt.Println("left", leftIdx)
		//fmt.Println("right", rightIdx)
	}


	//fmt.Println("final l:", leftIdx)
	//fmt.Println("final r:", rightIdx)

	if leftIdx == 0 && rightIdx == len(str) {
		//fmt.Println("whole")
		return str[leftIdx:rightIdx+1]
	}
	// if the palindrome found is a part of the section
	// return str[leftIdx+1:rightIdx]
	//fmt.Println("mid")
	return str[leftIdx+1:rightIdx]
}

func main() {
	// testString := "substring"
	// subString := testString[0:10]

	// fmt.Println(subString)

	test := "it's highnoon"

	//subResult := palindromeLengthCheck(test, 2, 3)
	//fmt.Println(subResult)

	wholeResult := LongestPalindromicSubstring(test)
	fmt.Println(wholeResult)

	// result := LongestPalindromicSubstring(test)
	// fmt.Println(result)
}