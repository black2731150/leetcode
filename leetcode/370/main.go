package main

import "fmt"

func getModifiedArray(length int, updates [][]int) []int {
	answer := make([]int, length)

	for _, v := range updates {
		start := v[0]
		end := v[1]
		num := v[2]

		for j := start; j <= end; j++ {
			answer[j] += num
		}
	}

	return answer
}

func main() {
	lenght := 5
	updates := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(getModifiedArray(lenght, updates))
}
