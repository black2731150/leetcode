package main

import (
	"fmt"
	"sort"
	"strings"
)

func beforeAndAfterPuzzles(phrases []string) []string {
	answer := map[string]struct{}{}
	headMap, endMap := make(map[string][]int), make(map[string][]int)

	for i, phrase := range phrases {
		p := strings.Fields(phrase)
		headMap[p[0]] = append(headMap[p[0]], i)
		endMap[p[len(p)-1]] = append(endMap[p[len(p)-1]], i)
	}

	for end, endIndexs := range endMap {
		if headIndex, find := headMap[end]; find {

			for _, v := range endIndexs {
				for _, v2 := range headIndex {
					if v != v2 {
						first := phrases[v]
						second := phrases[v2]
						f1 := strings.Fields(first)
						s1 := strings.Fields(second)
						a := append(f1, s1[1:]...)
						answer[strings.Join(a, " ")] = struct{}{}
					}
				}
			}
		}
	}

	a := []string{}
	for k := range answer {
		a = append(a, k)
	}

	sort.Strings(a)
	return a
}

func main() {
	phrases := []string{"a", "b", "a"}
	fmt.Println(beforeAndAfterPuzzles(phrases))
}
