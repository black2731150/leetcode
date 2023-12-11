package main

import "fmt"

type point struct {
	x int
	y int
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])

	maxHeght, minHight := heights[0][0], heights[0][0]
	for _, v := range heights {
		for _, v2 := range v {
			if v2 > maxHeght {
				maxHeght = v2
			}

			if v2 < minHight {
				minHight = v2
			}
		}
	}

	left, right := 0, abs(maxHeght-minHight)

	for left <= right {
		mid := (left + right) / 2

		hasGet := make([][]bool, n)
		for i := range hasGet {
			hasGet[i] = make([]bool, m)
		}

		queue := []point{{x: 0, y: 0}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if hasGet[p.x][p.y] {
				continue
			}
			hasGet[p.x][p.y] = true

			if p.x == n-1 && p.y == m-1 {
				right = mid - 1
				break
			}

			if p.x+1 < n && !hasGet[p.x+1][p.y] && abs(heights[p.x+1][p.y]-heights[p.x][p.y]) <= mid {
				queue = append(queue, point{p.x + 1, p.y})
			}

			if p.x-1 >= 0 && !hasGet[p.x-1][p.y] && abs(heights[p.x-1][p.y]-heights[p.x][p.y]) <= mid {
				queue = append(queue, point{p.x - 1, p.y})
			}

			if p.y+1 < m && !hasGet[p.x][p.y+1] && abs(heights[p.x][p.y+1]-heights[p.x][p.y]) <= mid {
				queue = append(queue, point{p.x, p.y + 1})
			}

			if p.y-1 >= 0 && !hasGet[p.x][p.y-1] && abs(heights[p.x][p.y-1]-heights[p.x][p.y]) <= mid {
				queue = append(queue, point{p.x, p.y - 1})
			}
		}

		if !hasGet[n-1][m-1] {
			left = mid + 1
		}
	}

	return left
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
}

//时间复杂度过高
// func minimumEffortPath(heights [][]int) int {
// 	n, m := len(heights), len(heights[0])

// 	realAnswer := 10000000
// 	RealHasGet := make([][]bool, n)
// 	for i := range RealHasGet {
// 		RealHasGet[i] = make([]bool, m)
// 	}

// 	var dfs func(lastHight, row, col, answer int)
// 	dfs = func(lastHight, row, col, answer int) {
// 		RealHasGet[row][col] = true
// 		answer = max(answer, abs(lastHight-heights[row][col]))

// 		if row == n-1 && col == m-1 {
// 			realAnswer = min(realAnswer, answer)
// 			RealHasGet[row][col] = false
// 			return
// 		}

// 		if row+1 < n && !RealHasGet[row+1][col] {
// 			dfs(heights[row][col], row+1, col, answer)
// 		}
// 		if row-1 >= 0 && !RealHasGet[row-1][col] {
// 			dfs(heights[row][col], row-1, col, answer)
// 		}
// 		if col+1 < m && !RealHasGet[row][col+1] {
// 			dfs(heights[row][col], row, col+1, answer)
// 		}
// 		if col-1 >= 0 && !RealHasGet[row][col-1] {
// 			dfs(heights[row][col], row, col-1, answer)
// 		}

// 		RealHasGet[row][col] = false
// 	}

// 	dfs(heights[0][0], 0, 0, 0)

// 	return realAnswer
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
