package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	answerSlice := [][]int{points[0]}

	for i := 0; i < len(points); i++ {
		if answerSlice[len(answerSlice)-1][1] >= points[i][0] {
			answerSlice[len(answerSlice)-1] = []int{min(answerSlice[len(answerSlice)-1][0], points[i][0]), min(answerSlice[len(answerSlice)-1][1], points[i][1])}
		} else {
			answerSlice = append(answerSlice, points[i])
		}
	}

	return len(answerSlice)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	points := [][]int{{1, 6}, {2, 8}, {7, 12}, {10, 16}}
	fmt.Println(findMinArrowShots(points))
}
