package main

import (
	"fmt"
	"sort"
)

func tow(times, scors []int, allTime int) int {
	Len := len(times)
	xjb := make([]float64, Len)
	theMap := make(map[float64]int, Len)
	for i := 0; i < Len; i++ {
		xjb[i] = float64(scors[i]) / float64(times[i])
		theMap[xjb[i]] = i
	}

	sort.Float64s(xjb)

	answer := 0
	for i := Len - 1; i >= 0; i-- {
		if allTime-times[theMap[xjb[i]]] > 0 {
			allTime = allTime - times[theMap[xjb[i]]]
			answer = answer + scors[theMap[xjb[i]]]
		}
	}

	return answer

}

func main() {

	var times []int
	fmt.Scanln(&times)
	var scores []int
	fmt.Scanln(&scores)
	t := 0
	fmt.Scan(&t)
	fmt.Println(tow(times, scores, t))
}
