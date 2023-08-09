package main

import (
	"fmt"
)

func nearestValidPoint(x int, y int, points [][]int) int {
	answer := -1
	min := 10000
	for i, point := range points {
		if x == point[0] || y == point[1] {
			if tmp := abs(x-point[0]) + abs(y-point[1]); tmp < min {
				min = tmp
				answer = i
			}
		}
	}
	return answer
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

func main() {
	x := 3
	y := 4
	points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	fmt.Println(nearestValidPoint(x, y, points))
}
