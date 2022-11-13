package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	array = append(array, n.Name) // creat the output array og node names in the graph
	for _, child := range n.Children { // for each chid of the of the current *Node, add its child to the output
		array = child.DepthFirstSearch(array)
	}
	return array
}


