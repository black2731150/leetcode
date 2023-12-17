package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	level := len(triangle)
	if level == 1 {
		return triangle[0][0]
	}

	answer := make([][]int, level)
	for i := range answer {
		answer[i] = make([]int, 0)
	}

	answer[0] = append(answer[0], triangle[0][0])
	for i := 1; i < level; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				answer[i] = append(answer[i], answer[i-1][0]+triangle[i][j])
				continue
			}

			if j == i {
				answer[i] = append(answer[i], answer[i-1][j-1]+triangle[i][j])
				continue
			}

			x := min(answer[i-1][j-1], answer[i-1][j]) + triangle[i][j]
			answer[i] = append(answer[i], x)
		}
	}

	realAnswer := answer[len(answer)-1][0]
	for _, v := range answer[len(answer)-1] {
		realAnswer = min(realAnswer, v)
	}
	return realAnswer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
