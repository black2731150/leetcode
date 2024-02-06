package main

import "fmt"

func findString(words []string, s string) int {
	n := len(words)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		m := mid
		for mid <= right && words[mid] == "" {
			mid++
		}

		if mid > right {
			right = m - 1
			continue
		}

		if words[mid] == s {
			return mid
		}

		if words[mid] > s {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	words := []string{"sd", "", "", "", "dsd"}
	s := "sd"
	fmt.Println(findString(words, s))
}
