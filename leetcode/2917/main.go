package main

import "fmt"

func findKOr(nums []int, k int) int {
	numsOfOne := make([]int, 32)

	answer := 0

	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if (num >> i) == 0 {
				break
			} else {
				if (num>>i)&1 == 1 {
					numsOfOne[i]++
				}
			}
		}
	}

	for i, n := range numsOfOne {
		if n >= k {
			answer = answer + 1<<i
		}
	}

	return answer
}

func main() {
	nums := []int{7, 12, 9, 8, 9, 15}
	k := 4
	fmt.Println(findKOr(nums, k))
}
