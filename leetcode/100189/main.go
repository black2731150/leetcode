package main

import "fmt"

func resultGrid(image [][]int, threshold int) [][]int {
	m := len(image)
	n := len(image[0])

	qiangdu := func(x, y int) int {
		if x-1 < 0 || x+1 >= m || y-1 < 0 || y+1 >= n {
			return -1
		}

		a1, b1, c1 := image[x-1][y-1], image[x-1][y], image[x-1][y+1]
		a2, b2, c2 := image[x][y-1], image[x][y], image[x][y+1]
		a3, b3, c3 := image[x+1][y-1], image[x+1][y], image[x+1][y+1]

		maxIntensity := 0
		maxIntensity = max(abs(a1-b1), maxIntensity)
		maxIntensity = max(abs(a2-b2), maxIntensity)
		maxIntensity = max(abs(a3-b3), maxIntensity)
		maxIntensity = max(abs(b1-c1), maxIntensity)
		maxIntensity = max(abs(b2-c2), maxIntensity)
		maxIntensity = max(abs(b3-c3), maxIntensity)
		maxIntensity = max(abs(a1-a2), maxIntensity)
		maxIntensity = max(abs(a2-a3), maxIntensity)
		maxIntensity = max(abs(b1-b2), maxIntensity)
		maxIntensity = max(abs(b2-b3), maxIntensity)
		maxIntensity = max(abs(c1-c2), maxIntensity)
		maxIntensity = max(abs(c2-c3), maxIntensity)

		if maxIntensity > threshold {
			return -1
		}

		return (a1 + a2 + a3 + b1 + b2 + b3 + c1 + c2 + c3) / 9
	}

	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	for i := range image {
		for j := range image[i] {
			sum := 0
			num := 0
			if x := qiangdu(i-1, j-1); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i-1, j); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i-1, j+1); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i, j-1); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i, j+1); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i, j); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i+1, j-1); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i+1, j); x != -1 {
				sum += x
				num++
			}
			if x := qiangdu(i+1, j+1); x != -1 {
				sum += x
				num++
			}

			if num > 0 {
				result[i][j] = sum / num
			} else {
				result[i][j] = image[i][j]
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	image := [][]int{{4, 8, 5}, {5, 7, 7}, {9, 4, 2}}
	t := 4
	fmt.Println(resultGrid(image, t))
}
