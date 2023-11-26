package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	fangxiang := true //true表示右上角移动，false表示左下角移动

	m := len(mat)
	n := len(mat[0])

	i := 0
	j := 0

	answer := make([]int, m*n)
	answerIndex := 0
	for i < m && j < n {

		// fmt.Printf("i=%d,j=%d\n", i, j)

		answer[answerIndex] = mat[i][j]
		answerIndex++

		if fangxiang { //如果右上角移动
			if i-1 < 0 || j+1 >= n { //如果超出界限了，就改变方向
				fangxiang = !fangxiang
				if j+1 < n {
					j = j + 1
				} else {
					i = i + 1
				}
			} else { //没有超出方向就移动到这个方向
				i, j = i-1, j+1
			}
		} else { //如果左下角移动
			if i+1 > m-1 || j-1 < 0 { //如果超出界限了，就改变方向
				fangxiang = !fangxiang
				if i+1 < m {
					i = i + 1
				} else {
					j = j + 1
				}
			} else { //如果没有超出界限，就移动到这个方向
				i, j = i+1, j-1
			}
		}
	}

	return answer
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(findDiagonalOrder(mat))
}
