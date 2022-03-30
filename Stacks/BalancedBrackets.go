package main

type Stack []rune

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str rune) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return 32, false // 32 is the rune value for a blank character
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isOdd(n string) bool {
	length := len(n)
	return length % 2 != 0
}

// Helper function to help determine if braket is opening or not
func isOpening(n byte) bool {
	bracket := string(n)
	return bracket == "(" || bracket == "[" || bracket == "{"
}

func isClosing(n byte) bool {
	bracket := string(n)
	return bracket == ")" || bracket == "]" || bracket == "}"
}

// Helper function to help determine if braket pairs are a match
func isMatch(a, b rune) bool {
	if string(a) == "(" && string(b) == ")" {
		return true
	} else if string(a) == "[" && string(b) == "]" {
		return true
	} else if string(a) == "{" && string(b) == "}" {
		return true
	}
	
	return false
}

func BalancedBrackets(s string) bool {
	// Write your code here.
	var stack Stack // create a stack variable of type Stack

	if isOdd(s) {
		return false
	}
	
	for _, bracket := range s {
		if isOpening(byte(bracket)) {
			stack.Push(bracket)
		} else if len(stack) > 0 && isMatch(stack[len(stack)-1], bracket) {
			stack.Pop()
		} else if len(stack) > 0 && !isMatch(stack[len(stack)-1], bracket) || isClosing(byte(bracket)){
			return false
		} else {
			continue
		}
	}

	if stack.IsEmpty() {
		return true
	}
	return false
}
