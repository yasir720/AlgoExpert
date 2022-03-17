package main

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	// Write your code here.
	queue := []*Node{n} // means that queue is a slice of pointers to Nodes of type Node
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		array = append(array, current.Name)

		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
	return array
}