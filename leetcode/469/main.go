package main

import "fmt"

func isConvex(points [][]int) bool {
	n := len(points)

	if n < 3 {
		return false
	}

	head, end := 0, 0
	for i := range points {
		firstPoint := points[i]
		secondPoint := points[(i+1)%n]
		thridPoint := points[(i+2)%n]

		x1, y1 := secondPoint[0]-firstPoint[0], secondPoint[1]-firstPoint[1]
		x2, y2 := thridPoint[0]-secondPoint[0], thridPoint[1]-secondPoint[1]

		end = x1*y2 - x2*y1
		if end != 0 {
			if head*end < 0 {
				return false
			}
			head = end
		}
	}
	return true
}

func main() {
	point := [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	fmt.Println(isConvex(point))
}
