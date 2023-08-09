package main

import "fmt"

func checkXMatrix(grid [][]int) bool {
	for i, v := range grid {
		for j, v1 := range v {
			if i == j || i+j == len(grid)-1 {
				if v1 == 0 {
					return false
				}
			} else {
				if v1 != 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	grid := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(checkXMatrix(grid))
}
