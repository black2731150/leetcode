package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)

	answer := make([]int, n)

	fullMap := make(map[int]int)

	heep := []int{}
	for i := range rains {
		answer[i] = 1
	}

	for i, v := range rains {
		if v > 0 {
			answer[i] = -1

			if day, find := fullMap[v]; find {
				index := sort.SearchInts(heep, day)
				if index == len(heep) {
					return []int{}
				}
				answer[heep[index]] = v
				heep = append(heep[:index], heep[index+1:]...)
			}
			fullMap[v] = i
		} else {
			heep = append(heep, i)
		}
	}

	return answer
}

func main() {
	rains := []int{1, 2, 3, 4}
	fmt.Println(avoidFlood(rains))
}
