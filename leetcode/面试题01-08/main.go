package main

func setZeroes(matrix [][]int) {
	hangMap, lieMap := make(map[int]struct{}), make(map[int]struct{})
	for i, v := range matrix {
		for i2, v2 := range v {
			if v2 == 0 {
				hangMap[i] = struct{}{}
				lieMap[i2] = struct{}{}
			}
		}
	}

	M := len(matrix)
	N := len(matrix[0])

	for k := range hangMap {
		for i := 0; i < N; i++ {
			matrix[k][i] = 0
		}
	}

	for k := range lieMap {
		for i := 0; i < M; i++ {
			matrix[i][k] = 0
		}
	}
	return
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
}
