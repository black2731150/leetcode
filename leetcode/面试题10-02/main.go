package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	theMap := make(map[string]int)
	answer := [][]string{}
	ai := 0

	for _, v := range strs {
		b := []byte(v)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		if index, find := theMap[string(b)]; find {
			answer[index] = append(answer[index], v)
		} else {
			theMap[string(b)] = ai
			ai++
			answer = append(answer, []string{v})
		}
	}
	return answer
}
