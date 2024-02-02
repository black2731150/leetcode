package main

import "fmt"

func oneEditAway(first string, second string) bool {
	if abs(len(first)-len(second)) > 1 {
		return false
	}

	if len(first) > len(second) {
		return oneEditAway(second, first)
	}

	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			continue
		} else {
			if isSameString(first[:i]+string(second[i])+first[i:], second) || isSameString(second[:i]+string(first[i])+second[i+1:], first) {
				return true
			} else if isSameString(first[:i]+first[i+1:], second) || isSameString(first, second[:i]+second[i+1:]) {
				return true
			} else if isSameString(first[:i]+string(second[i])+first[i+1:], second) || isSameString(second[:i]+string(first[i])+second[i+1:], first) {
				return true
			} else {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return 0
}

func isSameString(s1, s2 string) bool {
	return s1 == s2
}

func main() {
	frist := "pale"
	second := "ple"
	fmt.Println(oneEditAway(frist, second))
}
