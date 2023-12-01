package main

import "fmt"

func findCircleNum(isConnected [][]int) int {
	theMap := make(map[int][]int)

	n := 1
	for city1, v := range isConnected {
		city1++
		for city2, connected := range v {
			city2++

			if city1 > n {
				n = city1
			}
			if city2 > n {
				n = city2
			}
			if connected == 1 {
				theMap[city1] = append([]int{city2}, theMap[city1]...)
				theMap[city2] = append([]int{city1}, theMap[city2]...)
			}
		}

	}

	haveGet := make(map[int]struct{})

	answer := 0
	for i := 1; i <= n; i++ {

		if _, isFind := haveGet[i]; isFind {
			continue
		}

		answer++
		qurre := []int{i}
		for len(qurre) > 0 {
			copyQueue := make([]int, len(qurre))
			copy(copyQueue, qurre)
			qurre = []int{}
			for _, v := range copyQueue {
				citys := theMap[v]
				for _, v2 := range citys {
					if _, find := haveGet[v2]; find {
						continue
					} else {
						qurre = append(qurre, v2)
						haveGet[v2] = struct{}{}
					}
				}
			}
		}
	}

	return answer

}

func main() {
	isConnected := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	fmt.Println(findCircleNum(isConnected))
}
