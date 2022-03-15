package main

import (
	"fmt"
	"sort"
)

// func GroupAnagrams(words []string) [][]string {
// 	// Write your code here.
// 	return nil
// }

func main() {
	poop := map[string][]string{}
	poop["test"] = append(poop["test"], "hey")
	poop["test"] = append(poop["test"], "there")


	people := []string{"a", "c", "b"}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i] > people[j]
	})
	fmt.Println(people)
}