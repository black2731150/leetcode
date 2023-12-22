package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	theWordMap := make(map[string]struct{})
	for _, v := range wordDict {
		theWordMap[v] = struct{}{}
	}

	bp := make([]bool, len(s)+1)
	bp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if _, find := theWordMap[s[j:i]]; bp[j] && find {
				bp[i] = true
			}
		}
	}

	return bp[len(s)]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
