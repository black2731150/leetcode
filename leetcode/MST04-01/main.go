package main

import "fmt"

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	theMap := make(map[int][]int, n)

	for _, v := range graph {
		s, e := v[0], v[1]
		theMap[s] = append(theMap[s], e)
	}

	haveGet := make(map[int]struct{})

	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == target {
			return true
		}
		haveGet[node] = struct{}{}

		for _, v := range theMap[node] {
			if _, find := haveGet[v]; find {
				continue
			} else {
				queue = append(queue, v)
			}
		}
	}

	return false
}

func main() {
	n := 3
	graph := [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}
	start := 0
	end := 2
	fmt.Println(findWhetherExistsPath(n, graph, start, end))
}
