package main

// O(o + e) time | O(v) space
// e = number of edges & v = number of vertices in the graph
func TwoColorable(edges [][]int) bool {
	// Write your code here.
	colors := map[int]bool{ // use true/false for colors
		0: true,
	}
	stack := []int{0}

	for len(stack) > 0 {
		var node int
		node, stack = stack[len(stack)-1], stack[:len(stack)-1] // node is end stack, Stack end pop
		for _, connection := range edges[node] {
			if _, colorFound := colors[connection]; !colorFound {
				colors[connection] = !colors[node]
				stack = append(stack, connection)
			} else if colors[connection] == colors[node] {
				return false
			}
		}
	}
	return true
}
