package main

import "fmt"

func reachableNodes(n int, edges [][]int, restricted []int) int {

	nodesMap := make(map[int][]int)
	restrictedMap := make(map[int]struct{})
	haveGet := make(map[int]struct{})

	for _, v := range restricted {
		restrictedMap[v] = struct{}{}
	}

	for _, v := range edges {
		nodesMap[v[0]] = append(nodesMap[v[0]], v[1])
		nodesMap[v[1]] = append(nodesMap[v[1]], v[0])
	}

	answer := 0
	queue := []int{0}
	for len(queue) > 0 {
		copyQueue := make([]int, len(queue))
		copy(copyQueue, queue)
		queue = []int{}
		for _, node := range copyQueue {

			if _, isFind := haveGet[node]; isFind {
				continue
			}

			if _, isFind := restrictedMap[node]; !isFind {
				answer++
				haveGet[node] = struct{}{}
				queue = append(queue, nodesMap[node]...)
			}
		}
	}

	return answer
}

func main() {
	n := 7
	edge := [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}
	restricted := []int{4, 5}
	fmt.Println(reachableNodes(n, edge, restricted))
}
