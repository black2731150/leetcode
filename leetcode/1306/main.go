package main

import "fmt"

func canReach(arr []int, start int) bool {
	n := len(arr)

	theMap := make(map[int]struct{}, n)
	var dfs func(s int) bool
	dfs = func(s int) bool {

		if s < 0 || s >= n {
			return false
		}

		if _, find := theMap[s]; find {
			return false
		}

		if arr[s] == 0 {
			return true
		}
		theMap[s] = struct{}{}

		return dfs(s+arr[s]) || dfs(s-arr[s])
	}

	return dfs(start)
}

func main() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	start := 5
	fmt.Println(canReach(arr, start))
}
