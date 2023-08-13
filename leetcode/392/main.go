package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	if len(s) > len(t) {
		return false
	}

	sIndex := 0
	tIndex := 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}

		tIndex++
	}

	return sIndex == len(s)
}

func main() {
	s := "abc"
	t := "xabddcw"
	fmt.Println(isSubsequence(s, t))
}
