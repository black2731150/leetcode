package main

import "fmt"

func closetTarget(words []string, target string, startIndex int) int {
	answerIndexs := []int{}
	for i, v := range words {
		if v == target {
			answerIndexs = append(answerIndexs, i)
		}
	}
	answer := -1

	if len(answerIndexs) == 0 {
		return -1
	}

	answer = 100000

	for _, index := range answerIndexs {
		answer = min(answer, abs((startIndex-index+len(words))%len(words)))
		answer = min(answer, abs(index-startIndex+len(words))%len(words))
	}

	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func main() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "hello"
	startIndex := 1
	fmt.Println(closetTarget(words, target, startIndex))
}
