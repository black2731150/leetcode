package main

import "fmt"

func getMinDistance(nums []int, target int, start int) int {
	answer := 1000
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			for j := 0; j < len(nums); j++ {
				if tmp := abs(i - start); tmp < answer {
					answer = tmp
				}
			}
		}
	}
	return answer
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	start := 3
	fmt.Println(getMinDistance(nums, target, start))
}
