package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	min := 201

	first := strs[0]
	if strsLen == 1 {
		return first
	}

	if first == "" {
		return first
	}

	if len(first) < min {
		min = len(first)
	}

	index := 0
	for i := 1; i < strsLen; i++ {
		if strs[i] == "" {
			break
		}

		if len(strs[i]) < min {
			min = len(strs[i])
		}

		if index >= min {
			break
		}

		if index < len(strs[i]) && strs[i][index] == first[index] {
			if i == strsLen-1 {
				index++
				i = 0
			}
		} else {
			break
		}
	}

	return first[:index]
}

func main() {
	strs := []string{"a", "ac"}
	fmt.Println(longestCommonPrefix(strs))
}
