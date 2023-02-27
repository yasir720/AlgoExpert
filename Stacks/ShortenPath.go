package main

import "strings"

func ShortenPath(path string) string {
	// Write your code here.
	isAbsPath := path[0] == '/'
	splits := strings.Split(path, "/")
	tokens := []string{}

	for _, token := range splits {
		if isImportantToken(token) {
			tokens = append(tokens, token)
		}
	}

	stack := []string{}
	if isAbsPath {
		stack = append(stack, "")
	}

	for _, token := range tokens {
		if token == ".." {
			if len(stack) == 0 || stack[len(stack)-1] == ".." {
				stack= append(stack, token)
			} else if stack[len(stack)-1] != "" {
				stack = stack[:len(stack)-1]
			}

		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) == 1 && stack [0] == "" {
		return "/"
	}
	
	return strings.Join(stack, "/")
}

func isImportantToken(token string) bool {
	return len(token) > 0 && token != "."
}
