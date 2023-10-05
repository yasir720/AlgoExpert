package main

func BestDigits(number string, numDigits int) string {
	// Write your code here.
	stack := make([]rune, 0)

	for _, digit :=  range number {
		for numDigits > 0 && len(stack) > 0 && digit > stack[len(stack)-1] {
			numDigits -= 1
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, digit)
	}

	for numDigits > 0 {
		numDigits -= 1
		stack = stack[:len(stack)-1]
	}
	return string(stack)
}
