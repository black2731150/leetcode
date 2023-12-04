package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	n := len(s)
	answer := [][]int{}
	if n < 3 {
		return answer
	}

	last := 0

	for i := 0; i < n; i++ {
		if s[i] == s[last] {
			continue
		} else {
			if i-last > 2 {
				answer = append(answer, []int{last, i - 1})
				last = i
			} else {
				last = i
			}
		}
	}

	if last != n-1 && n-last > 2 {
		answer = append(answer, []int{last, n - 1})
	}

	return answer
}

func main() {
	s := "abcdddeeeeaabbbcd"
	fmt.Println(largeGroupPositions(s))
}
