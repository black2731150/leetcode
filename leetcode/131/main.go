package main

import "fmt"

func partition(s string) [][]string {
	answer := [][]string{}

	n := len(s)

	bp := make([][]bool, n)
	for i := range bp {
		bp[i] = make([]bool, n)
		for j := range bp[i] {
			bp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			bp[i][j] = s[i] == s[j] && bp[i+1][j-1]
		}
	}

	tmp := []string{}
	var dfs func(index int)
	dfs = func(index int) {
		if index == n {
			answer = append(answer, append([]string{}, tmp...))
			return
		}

		for j := index; j < n; j++ {
			if bp[index][j] {
				tmp = append(tmp, s[index:j+1])
				dfs(j + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return answer
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
