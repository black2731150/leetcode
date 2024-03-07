package main

import "fmt"

func getLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)

	firstZero := 0
	for ; firstZero < n && groups[firstZero] != 0; firstZero++ {
	}

	var zero, one []int
	if firstZero != n {
		zero = help(groups[firstZero:])
	}

	firstOne := 0
	for ; firstOne < n && groups[firstOne] != 1; firstOne++ {
	}

	if firstOne != n {
		one = help(groups[firstOne:])
	}

	fmt.Println(zero, one)

	answer := []string{}
	if len(one) > len(zero) {
		for _, i := range one {
			answer = append(answer, words[firstOne+i])
		}
	} else {
		for _, i := range zero {
			answer = append(answer, words[firstZero+i])
		}
	}

	return answer
}

func help(groups []int) []int {
	answer := []int{}
	answer = append(answer, 0)
	last := 0
	for i := 1; i < len(groups); i++ {
		if groups[last] != groups[i] {
			last = i
			answer = append(answer, i)
		}
	}

	return answer
}

func main() {
	words := []string{"o", "cfy", "en"}
	groups := []int{1, 0, 0}
	fmt.Println(getLongestSubsequence(words, groups))
}
