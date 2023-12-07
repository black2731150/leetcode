package main

import "fmt"

func minReorder(n int, connections [][]int) int {
	FromMeToYou := make(map[int][]int)
	FromYouToMe := make(map[int][]int)
	for _, v := range connections {
		a := v[0]
		b := v[1]
		FromMeToYou[a] = append(FromMeToYou[a], b)
		FromYouToMe[b] = append(FromYouToMe[b], a)
	}

	canGetZeroMap := make(map[int]struct{})
	canGetZeroMap[0] = struct{}{}

	answer := 0
	var dfs func(n int) int
	dfs = func(n int) int {
		to, findTo := FromMeToYou[n]
		from, findFrom := FromYouToMe[n]

		if findTo {
			for _, v := range to {
				if _, canGetZero := canGetZeroMap[v]; !canGetZero {
					canGetZeroMap[v] = struct{}{}
					answer += dfs(v) + 1
				}
			}
		}

		if findFrom {
			for _, v := range from {
				if _, canGetZero := canGetZeroMap[n]; canGetZero {
					canGetZeroMap[v] = struct{}{}
					answer += dfs(v)
				}
			}
		}

		return 0
	}

	dfs(0)

	return answer
}

func main() {
	n := 5
	connections := [][]int{{4, 3}, {2, 3}, {1, 2}, {1, 0}}
	fmt.Println(minReorder(n, connections))
}
