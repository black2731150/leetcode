package main

import (
	"fmt"
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	nums := []int{}
	numsbool := []bool{}
	for _, v := range mat {
		x := 0
		for _, v1 := range v {
			if v1 == 1 {
				x++
			}
		}
		nums = append(nums, x)
		numsbool = append(numsbool, true)
	}

	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)

	sort.Ints(numsCopy)
	fmt.Println(nums)
	fmt.Println(numsCopy)

	answer := []int{}
	for i := 0; i < k; i++ {
		for j, v := range nums {
			if v == numsCopy[i] && numsbool[j] {
				answer = append(answer, j)
				numsbool[j] = false
				break
			}
		}
	}

	return answer

}

func main() {
	mat := [][]int{}
	mat = append(mat, []int{1, 1, 0, 0, 0})
	mat = append(mat, []int{1, 1, 1, 1, 0})
	mat = append(mat, []int{1, 0, 0, 0, 0})
	mat = append(mat, []int{1, 1, 0, 0, 0})
	mat = append(mat, []int{1, 1, 1, 1, 1})

	fmt.Println(kWeakestRows(mat, 3))
}
