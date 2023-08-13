package main

import "fmt"

func strStr(haystack string, needle string) int {
	first := needle[0]

	for i := range haystack {
		if haystack[i] == first {
			index := 0
			for j := i; j < len(haystack) && index < len(needle); {
				if haystack[j] == needle[index] {

					if index == len(needle)-1 {
						return i
					}

					index++
					j++
					continue
				} else {
					break
				}
			}
		}
	}

	return -1
}

func main() {
	haystack := "leetcode"
	needle := "leeto"
	fmt.Println(strStr(haystack, needle))
}
