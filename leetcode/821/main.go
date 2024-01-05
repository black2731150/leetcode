package main

import "fmt"

func shortestToChar(s string, c byte) []int {
	indexs := []int{}

	for i := range s {
		if s[i] == c {
			indexs = append(indexs, i)
		}
	}

	answer := make([]int, len(s))
	for i := range answer {
		one := len(s)
		for _, v := range indexs {
			one = min(one, abs(v-i))
		}
		answer[i] = one
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := "loveleetcode"
	c := 'e'
	fmt.Println(shortestToChar(s, byte(c)))
}
