package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	answer := [][]string{}

	sortedStrs := make([]string, len(strs))
	for i := range strs {
		sortedStrs[i] = sortString(strs[i])
	}

	theMap := make(map[string][]int)
	for i := range sortedStrs {
		theMap[sortedStrs[i]] = append(theMap[sortedStrs[i]], i)
	}

	for _, v := range theMap {
		x := []string{}
		for i := range v {
			x = append(x, strs[v[i]])
		}
		answer = append(answer, x)
	}
	return answer
}

func sortString(s string) string {
	bytes := make([]int, len(s))
	for i := range s {
		bytes[i] = int(s[i])
	}

	sort.Ints(bytes)

	answer := ""
	for i := range bytes {
		answer = answer + string(rune(bytes[i]))
	}
	return answer
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
