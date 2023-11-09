package main

import "fmt"

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		numsA := make([]int, n)
		numsB := make([]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j && knows(i, j) {
					fmt.Println(i, " konws ", j)
					numsA[i]++
					numsB[j]++
				}
			}
		}

		answer := -1
		for i, v := range numsB {
			if v == n-1 {
				if numsA[i] == 0 {
					answer = i
				}
			}
		}

		return answer

	}
}

func knows(a, b int) bool {
	return a > b
}

func main() {
	solution(knows)
}
