package main

import "fmt"

func numberOfBoomerangs(points [][]int) int {

	answer := 0
	for i := 0; i < len(points); i++ {
		theMap := make(map[int]int)
		for j := 0; j < len(points); j++ {
			theMap[help(points[i][0], points[i][1], points[j][0], points[j][1])]++
		}
		for _, v := range theMap {
			answer += pailie(v)
		}
	}

	return answer
}

func pailie(n int) int {
	return n * (n - 1)
}

func help(xi, yi, xj, yj int) int {
	return abs(squar(xi-xj) + squar(yi-yj))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func squar(x int) int {
	return x * x
}

func main() {
	points := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println(numberOfBoomerangs(points))
}
