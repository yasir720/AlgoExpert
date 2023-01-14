package main

// O(n^2) time | O(n) space - where n is the length of the stack
func SortStack(stack []int) []int {
	// Write your code here.
	if len(stack) == 0 {
		return stack
	}

	// we "pop" the stack
	top := stack[len(stack)-1] // get the top value
	stack = stack[:len(stack)-1] // cut off the top element

	SortStack(stack)

	insertInSortedOrder(&stack, top)
	return stack
}

func insertInSortedOrder(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value)
		return
	}

	// we "pop" the stack
	top := (*stack)[len(*stack)-1] // get the top value
	*stack = (*stack)[:len(*stack)-1] // cut off the top element
	
	insertInSortedOrder(stack, value)
	*stack = append(*stack, top)

}