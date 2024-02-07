package main

import "fmt"

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	r := len(obstacleGrid) - 1
	c := len(obstacleGrid[0]) - 1
	answer := [][]int{}

	var dfs func(path [][]int)
	dfs = func(path [][]int) {
		if len(answer) == 0 {
			i, j := path[len(path)-1][0], path[len(path)-1][1]
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = 1
				if i < r {
					dfs(append(path, []int{i + 1, j}))
				}

				if j < c {
					dfs(append(path, []int{i, j + 1}))
				}

				if i == r && j == c {
					answer = make([][]int, len(path))
					copy(answer, path)
				}
			}
		}
	}
	dfs([][]int{{0, 0}})

	return answer
}

func main() {
	obstacleCrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(pathWithObstacles(obstacleCrid))
}
