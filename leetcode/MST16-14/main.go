package main

import (
	"math"
)

func bestLine(points [][]int) []int {

	maxNum := 0
	answer := []int{}

	n := len(points)
	for i := 0; i < n; i++ {
		LineMap := make(map[float64][]int)
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]

			if x2 == x1 {
				LineMap[math.Inf(1)] = append(LineMap[math.Inf(1)], j)
			} else {
				k := (float64(y2) - float64(y1)) / (float64(x2) - float64(x1))
				LineMap[k] = append(LineMap[k], j)
			}
		}

		for _, v := range LineMap {
			if len(v)+1 > maxNum {
				answer = append([]int{i}, v...)
				maxNum = len(v) + 1
			}
		}
	}

	return answer[:2]
}

func main() {
	points := [][]int{{0, 0}, {0, 0}, {1, 0}, {2, 0}}
	bestLine(points)
}
