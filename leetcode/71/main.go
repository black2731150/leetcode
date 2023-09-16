package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	if path[0] != '/' {
		return ""
	}

	peases := strings.Split(path, "/")

	pathStack := []string{}

	for _, pease := range peases {

		if pease == "" {
			continue
		}

		if pease == "." {
			continue
		}

		if pease == ".." {
			if len(pathStack) != 0 {
				pathStack = pathStack[:len(pathStack)-1]
			}
			continue
		}

		pathStack = append(pathStack, pease)
	}

	if len(pathStack) == 0 {
		return "/"
	}

	answer := ""
	for i := 0; i < len(pathStack); i++ {
		answer = answer + "/" + pathStack[i]
	}

	return answer
}

func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}
