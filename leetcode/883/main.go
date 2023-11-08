package main

import "fmt"

func projectionArea(grid [][]int) int {
	xy := 0
	yz := 0
	xz := 0

	n := len(grid)

	yzMaxs := make([]int, n)
	xzMaxs := make([]int, n)

	for i, v := range grid {
		for i2, v2 := range v {
			if v2 > 0 {
				xy++
			}
			if v2 > yzMaxs[i] {
				yzMaxs[i] = v2
			}

			if v2 > xzMaxs[i2] {
				xzMaxs[i2] = v2
			}
		}
	}

	for _, v := range yzMaxs {
		yz += v
	}

	for _, v := range xzMaxs {
		xz += v
	}

	return xy + yz + xz
}

func main() {
	gird := [][]int{{1, 2}, {1, 2}}
	fmt.Println(projectionArea(gird))
}
