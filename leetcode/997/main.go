package main

import "fmt"

func findJudge(n int, trust [][]int) int {
	themap := map[int][]int{}
	theTrustNumMap := make([]int, n)
	for i := 1; i <= n; i++ {
		themap[i] = []int{}
		theTrustNumMap[i-1] = 0
	}

	for _, v := range trust {
		themap[v[0]] = append(themap[v[0]], v[1])
		theTrustNumMap[v[1]-1]++
	}
	fmt.Println(themap)
	fmt.Println(theTrustNumMap)
	for i := range themap {
		if len(themap[i]) == 0 {
			if theTrustNumMap[i-1] == n-1 {
				return i
			}
		}
	}

	return -1
}

func main() {
	n := 3
	trust := [][]int{}
	trust = append(trust, []int{1, 2}, []int{2, 3})
	fmt.Println(findJudge(n, trust))
}
