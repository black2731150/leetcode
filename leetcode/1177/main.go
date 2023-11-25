package main

import "fmt"

func canMakePaliQueries(s string, queries [][]int) []bool {

	n := len(s)
	prefixSum := make([][26]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i]
		prefixSum[i+1][s[i]-'a']++
	}

	answer := make([]bool, len(queries))
	for i, v := range queries {
		left := v[0]
		right := v[1]
		k := v[2]
		oddnum := 0
		for j := 0; j < 26; j++ {
			count := prefixSum[right+1][j] - prefixSum[left][j]
			if count%2 == 1 {
				oddnum++
			}
		}
		oddnum = max(oddnum, oddnum-1)
		answer[i] = oddnum/2 <= k
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "abcda"
	queries := [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
	fmt.Println(canMakePaliQueries(s, queries))
}
