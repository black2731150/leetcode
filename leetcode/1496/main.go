package main

import "fmt"

type point struct {
	x int
	y int
}

func isPathCrossing(path string) bool {
	theMap := make(map[point]bool)

	i, j := 0, 0
	theMap[point{x: 0, y: 0}] = true
	for _, p := range path {
		if p == 'W' {
			i++
		}

		if p == 'S' {
			j--
		}

		if p == 'N' {
			j++
		}

		if p == 'E' {
			i--
		}

		if _, find := theMap[point{x: i, y: j}]; find {
			return true
		} else {
			theMap[point{x: i, y: j}] = true
		}
	}

	return false
}

func main() {
	path := "NES"
	fmt.Println(isPathCrossing(path))
}
