package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); {
		if intervals[i][0] >= intervals[i-1][0] && intervals[i][1] <= intervals[i-1][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}
	return len(intervals)
}

func main() {
	intervals := [][]int{{1, 2}, {1, 4}, {3, 4}}
	fmt.Println(removeCoveredIntervals(intervals))
}
