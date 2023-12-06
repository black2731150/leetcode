package main

import (
	"fmt"
	"sort"
)

func pancakeSort(arr []int) []int {
	answer := []int{}
	for !sort.IntsAreSorted(arr) {
		mIndex := maxIndex(arr)
		if mIndex == -1 {
			break
		}

		if mIndex == len(arr)-1 {
			arr = arr[:len(arr)-1]
			continue
		}

		if mIndex == 0 {
			for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
				arr[i], arr[j] = arr[j], arr[i]
			}
			answer = append(answer, len(arr))
			continue
		}

		for i, j := 0, len(arr[:mIndex+1])-1; i < j; i, j = i+1, j-1 {
			arr[:mIndex+1][i], arr[:mIndex+1][j] = arr[:mIndex+1][j], arr[:mIndex+1][i]
		}
		answer = append(answer, mIndex+1)
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		answer = append(answer, len(arr))
	}

	return answer
}

func maxIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	max := 0
	for i, v := range arr {
		if v > arr[max] {
			max = i
		}
	}

	return max
}

func main() {
	arr := []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(arr))
}
