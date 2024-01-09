package main

import (
	"fmt"
)

type line struct {
	k          float64
	b          float64
	isVertical bool
}

func maxPoints(points [][]int) int {
	maxNum := 0

	for i := 0; i < len(points); i++ {
		localMax := 0
		duplicate := 1
		theMap := make(map[line]int)

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				duplicate++
				continue
			}

			l := GetLine(points[i][0], points[i][1], points[j][0], points[j][1])
			theMap[l]++
			localMax = max(localMax, theMap[l])
		}

		maxNum = max(maxNum, localMax+duplicate)
	}

	return maxNum
}

func GetLine(x1, y1, x2, y2 int) line {
	answer := line{}
	if y1 == y2 {
		answer.k = 0
		answer.b = float64(y1)
		return answer
	}

	if x1 == x2 {
		answer.k = 0
		answer.b = float64(x1)
		answer.isVertical = true
		return answer
	}

	answer.k = float64(y2-y1) / float64(x2-x1)
	answer.b = float64(y1) - answer.k*float64(x1)
	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{0, 0}, {0, 1}}
	fmt.Println(maxPoints(points))
}
