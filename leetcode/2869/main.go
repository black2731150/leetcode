package main

import "fmt"

func minOperations(nums []int, k int) int {
	haveNum := make([]bool, k)
	numSum := 0
	answer := -1
	for i := len(nums) - 1; i >= 0; i-- {
		v := nums[i]
		if v <= k {
			if !haveNum[v-1] {
				numSum++
				haveNum[v-1] = true
				if numSum == k {
					answer = len(nums) - i
					break
				}
			}
		}
	}
	return answer
}

func main() {
	nums := []int{2, 3, 1, 4, 7, 8}
	k := 4

	fmt.Println(minOperations(nums, k))
}
