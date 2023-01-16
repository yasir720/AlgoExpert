package main

import "fmt"

func main() {
	stack := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(stack[len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	// print versions
	fmt.Print()
	fmt.Printf()
	fmt.Println()


}
