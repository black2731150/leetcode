package main

import "fmt"

func minimumFuelCost(roads [][]int, seats int) int64 {
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var res int64
	var dfs func(int, int) int
	dfs = func(cur, fa int) int {
		peopleSum := 1
		for _, ne := range g[cur] {
			if ne != fa {
				peopleCnt := dfs(ne, cur)
				peopleSum, res = peopleSum+peopleCnt, res+int64((peopleCnt+seats-1)/seats)
			}
		}
		return peopleSum
	}
	dfs(0, -1)
	return res
}

func main() {
	roads := [][]int{{0, 1}, {0, 2}, {0, 3}}
	seats := 5
	fmt.Println(minimumFuelCost(roads, seats))

}
