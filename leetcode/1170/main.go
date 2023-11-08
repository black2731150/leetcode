package main

import (
	"fmt"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsNumSlice := make([]int, len(words))
	for i, s := range words {
		wordsNumSlice[i] = f(s)
	}

	answer := []int{}
	for _, word := range queries {
		x := f(word)
		ans := 0
		for _, v := range wordsNumSlice {
			if x < v {
				ans++
			}
		}
		answer = append(answer, ans)
	}

	return answer
}

func f(s string) int {
	answer := make([]int, 26)

	for i := range s {
		answer[s[i]-'a']++
	}

	for _, v := range answer {
		if v > 0 {
			return v
		}
	}
	return 0
}

func main() {
	s := "dcce"
	fmt.Println(f(s))
	q := []string{"sdsds", "sdsdsfdfdv", "dwiuudhniu"}
	words := []string{"a", "aaa", "aaaa", "ccccccccc"}
	fmt.Println(numSmallerByFrequency(q, words))
}
