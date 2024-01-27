package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return r == '.' || r == '!' || r == '?' || r == '\'' || r == '"' || r == ',' || r == ' ' || r == ';'
	})

	theBanMap := make(map[string]struct{})
	for _, v := range banned {
		theBanMap[v] = struct{}{}
	}

	theMap := make(map[string]int)

	answer := ""
	maxNum := 0
	for _, v := range words {
		if _, find := theBanMap[v]; find {
			continue
		}

		theMap[v]++
		if x := theMap[v]; x > maxNum {
			maxNum = x
			answer = v
		}

	}
	return answer
}

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}
