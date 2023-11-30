package main

import "fmt"

func partitionString(s string) int {
	charMap := make([]int, 26)
	answer := 0
	for _, ch := range s {
		if charMap[ch-'a'] == 0 {
			charMap[ch-'a']++
		} else {
			answer++
			charMap = make([]int, 26)
			charMap[ch-'a']++
		}
	}

	for _, v := range charMap {
		if v != 0 {
			answer++
			break
		}
	}

	return answer
}

func main() {
	s := "abcckscmksm"
	fmt.Println(partitionString(s))
}
