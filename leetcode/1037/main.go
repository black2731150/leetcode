package main

import (
	"fmt"
	"math"
)

func isBoomerang(points [][]int) bool {
	x1, y1 := points[0][0], points[0][1]
	x2, y2 := points[1][0], points[1][1]
	x3, y3 := points[2][0], points[2][1]

	if x1 == x2 && y1 == y2 {
		return false
	}

	if x1 == x3 && y1 == y3 {
		return false
	}

	if x3 == x2 && y3 == y2 {
		return false
	}

	var k1, k2 float64 = math.Inf(0), math.Inf(0)
	if x2-x1 != 0 {
		k1 = float64(y2-y1) / float64(x2-x1)
	}

	if x3-x2 != 0 {
		k2 = float64(y3-y2) / float64(x3-x2)
	}

	fmt.Println(k1, k2)

	return k1 != k2
}

func main() {
	points := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
	}

	fmt.Println(isBoomerang(points))
}
