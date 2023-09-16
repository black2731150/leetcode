package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {

	n := len(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	answer := [][]int{intervals[0]}

	for i := 1; i < n; i++ {
		if answer[len(answer)-1][1] >= intervals[i][0] {
			answer[len(answer)-1] = []int{answer[len(answer)-1][0], max(intervals[i][1], answer[len(answer)-1][1])}
		} else {
			answer = append(answer, intervals[i])
		}
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
	intervals := [][]int{
		{2, 6}, {1, 3}, {15, 18}, {8, 10},
	}

	fmt.Println(insert(intervals, []int{5, 7}))
}
