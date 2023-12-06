package main

import "fmt"

func numSpecial(mat [][]int) int {
	rowSum := make([]int, len(mat))
	colSum := make([]int, len(mat[0]))

	// 计算行和列的和
	for i := range mat {
		for j, value := range mat[i] {
			rowSum[i] += value
			colSum[j] += value
		}
	}

	answer := 0
	// 检查特殊元素
	for i := range mat {
		for j, value := range mat[i] {
			if value == 1 && rowSum[i] == 1 && colSum[j] == 1 {
				answer++
			}
		}
	}

	return answer
}

func main() {
	may := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Println(numSpecial(may))
}
