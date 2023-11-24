package main

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
	answer := []int{}
	for _, v := range queries {
		x := v[0]
		y := v[1]
		r := v[2]

		oneAnswer := 0

		for _, v2 := range points {
			ponitX := v2[0]
			pointY := v2[1]
			if (ponitX-x)*(ponitX-x)+(pointY-y)*(pointY-y) <= r*r {
				oneAnswer++
			}
		}
		answer = append(answer, oneAnswer)
	}

	return answer
}

func main() {
	points := [][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}
	queries := [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}
	fmt.Println(points, queries)
}
