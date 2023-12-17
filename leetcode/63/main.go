package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	answers := make([][]int, m)
	for i := range answers {
		answers[i] = make([]int, n)
	}

	answers[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				if obstacleGrid[i][j] == 1 {
					answers[i][j] = 0
				} else {
					if i == 0 && j > 0 {
						answers[i][j] = answers[i][j-1]
					}

					if j == 0 && i > 0 {
						answers[i][j] = answers[i-1][j]
					}
				}
			} else {
				if obstacleGrid[i][j] == 1 {
					answers[i][j] = 0
				} else {
					answers[i][j] = answers[i-1][j] + answers[i][j-1]
				}
			}
		}
	}

	return answers[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
