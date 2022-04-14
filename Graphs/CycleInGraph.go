package main

// O(v+e) time | O(V) space - v= # of total vertices - e = # of edges
func CycleInGraph(edges [][]int) bool {
	// Write your code here.
	numberOfNodes := len(edges)
	visited := make([]bool, len(edges))
	currentlyInStack := make([]bool, len(edges))

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
		} else if currentlyInStack[neighbor] {
			return true
		}
	}

	currentlyInStack[node] = false
	return false
}
