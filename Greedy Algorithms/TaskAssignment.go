package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(n) space
func TaskAssignment(k int, tasks []int) [][]int {
	// Write your code here.
	// 2D slice to hold assigned task pairs
	assignedTasks := make([][]int, 0)
	// map tasks to their inital index
	tasks2Idx := make(map[int][]int)
	for idx, taskDuration := range tasks {
		tasks2Idx[taskDuration] = append(tasks2Idx[taskDuration], idx)
	}

	// sort task duration to get optimal pairings
	sort.Ints(tasks)

	for i := 0; i < k; i++ {
		// select the best durations possible for each task
		task1 := tasks2Idx[tasks[i]][0]
		task2 := tasks2Idx[tasks[len(tasks)-(i+1)]][0]

		// assign optimal taks
		assignedTasks= append(assignedTasks, []int{task1, task2})

		// remove duplicate duration idexes if any
		tasks2Idx[tasks[i]] = tasks2Idx[tasks[i]][1:]
		tasks2Idx[tasks[len(tasks)-(i+1)]] = tasks2Idx[tasks[len(tasks)-(i+1)]][1:]	
	}

	return assignedTasks
}

func main() {
	test := []int{1, 3, 5, 3, 1, 4}

	result := TaskAssignment(3, test)

	fmt.Println(result)
}
