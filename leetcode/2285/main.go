package main

import (
	"fmt"
	"sort"
)

type City struct {
	cityNum int
	roadNum int
}

func maximumImportance(n int, roads [][]int) int64 {
	allCitys := make([]City, n)
	for i := 0; i < n; i++ {
		allCitys[i].cityNum = i
	}

	for _, v := range roads {
		city1 := v[0]
		city2 := v[1]
		allCitys[city1].roadNum++
		allCitys[city2].roadNum++
	}

	sort.Slice(allCitys, func(i, j int) bool {
		return allCitys[i].roadNum > allCitys[j].roadNum
	})

	var answer int64 = 0
	for _, c := range allCitys {
		answer = answer + int64(c.roadNum*n)
		n--
	}

	return answer
}

func main() {
	n := 5
	roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println(maximumImportance(n, roads))
}
