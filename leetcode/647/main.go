package main

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	bp := make([][]bool, n)
	for i := range bp {
		bp[i] = make([]bool, n)
	}

	answer := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 3 || bp[i+1][j-1]) {
				bp[i][j] = true
				answer++
			}
		}
	}

	return answer
}

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}
