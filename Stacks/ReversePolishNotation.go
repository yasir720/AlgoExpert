package main

import "strconv"

func ReversePolishNotation(tokens []string) int {
	// Write your code here.
	stack := make([]int, 0)

	for _, token := range tokens {
		if token ==  "+" {
			result := popStack(&stack) + popStack(&stack)
			stack = append(stack, result)
		} else if token ==  "-" {
			firstNum := popStack(&stack)
			result := popStack(&stack) - firstNum
			stack = append(stack, result)
		} else if token ==  "*" {
			result := popStack(&stack) * popStack(&stack)
			stack = append(stack, result)
		} else if token ==  "/" {
			firstNum := popStack(&stack)
			result := popStack(&stack) / firstNum
			stack = append(stack, result)
		} else {
			n, _ := strconv.Atoi(token)
			stack = append(stack, n)
		}
	}
	return stack[len(stack)-1]
}

func popStack(stack *[]int) int {
	var n int
	n, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
	return n
}
