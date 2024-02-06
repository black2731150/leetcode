package main

import "fmt"

func cutSquares(square1 []int, square2 []int) []float64 {
	c1x := float64(square1[0]) + float64(square1[2])/2.0
	c1y := float64(square1[1]) + float64(square1[2])/2.0
	c2x := float64(square2[0]) + float64(square2[2])/2.0
	c2y := float64(square2[1]) + float64(square2[2])/2.0
	if c1x == c2x {
		// 垂直分割
		return []float64{
			c1x,
			min(square1[1], square2[1]),
			c1x,
			max(square1[1]+square1[2], square2[1]+square2[2]),
		}
	} else {
		k := (c1y - c2y) / (c1x - c2x)
		b := c1y - k*c1x
		if abs(k) > 1 {
			// 底部
			bottomY := min(square2[1], square1[1])
			bottomX := (bottomY - b) / k
			// 顶部
			topY := max(square1[1]+square1[2], square2[1]+square2[2])
			topX := (topY - b) / k

			if topX < bottomX {
				return []float64{topX, topY, bottomX, bottomY}
			} else {
				return []float64{bottomX, bottomY, topX, topY}
			}
		} else {
			// 左边
			leftX := min(square1[0], square2[0])
			leftY := k*leftX + b

			// 右边
			rightX := max(square1[0]+square1[2], square2[0]+square2[2])
			rightY := k*rightX + b
			if leftX < rightX {
				return []float64{leftX, leftY, rightX, rightY}
			} else {
				return []float64{rightX, rightY, leftX, leftY}
			}
		}
	}
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}

func min(a int, b int) float64 {
	if a > b {
		return float64(b)
	}
	return float64(a)
}

func max(a int, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

func main() {
	square1 := []int{-1, 1, 5}
	square2 := []int{-2, -2, 6}
	fmt.Println(cutSquares(square1, square2))
}
