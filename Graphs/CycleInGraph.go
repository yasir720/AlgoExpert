package main

// O(v+e) time | O(V) space - v= # of total vertices - e = # of edges
func CycleInGraph(edges [][]int) bool {
	// Write your code here.
	numberOfNodes := len(edges)
	visited := make([]bool, len(edges)) // to help us keep track of the visited nodes
	currentlyInStack := make([]bool, len(edges)) // keep track of the recursive call - ancestors of nodes

	for node := 0; node < numberOfNodes; node++ {
		if visited[node] {
			continue
		}

		containsCycle := isNodeInCycle(node, edges, visited, currentlyInStack)
		if containsCycle {
			return true
		}
	}
	return false
}

func isNodeInCycle(node int, edges [][]int, visited, currentlyInStack []bool) bool {
	visited[node] = true
	currentlyInStack[node] = true

	neighbors := edges[node]
	for _, neighbor := range neighbors {
		if visited[neighbor] == false {
			ccontainsCycle := isNodeInCycle(neighbor, edges, visited, currentlyInStack)
			if ccontainsCycle {
				return true
			}
		} else if currentlyInStack[neighbor] { // and if the node has been visited. The case where we have found a cycle
			return true
		}
	}

	currentlyInStack[node] = false // to take the node off the call stack
	return false
}
