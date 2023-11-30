package main

import "fmt"

func getMaxLen(nums []int) int {
	maxLen, currPositive, currNegative := 0, 0, 0

	for _, num := range nums {
		if num > 0 {
			currPositive++
			if currNegative > 0 {
				currNegative++
			}
		} else if num < 0 {
			temp := currPositive
			if currNegative > 0 {
				currPositive = currNegative + 1
			} else {
				currPositive = 0
			}
			currNegative = temp + 1
		} else {
			currPositive, currNegative = 0, 0
		}
		maxLen = max(maxLen, currPositive)
	}

	return maxLen
}

//时间复杂度为O(n^2) 超时
// func getMaxLen(nums []int) int {
// 	sum := make([][2]int, len(nums)+1)
// 	sum[0] = [2]int{0, 0}

// 	answer := 0
// 	for i, v := range nums {
// 		lastZheng := sum[i][0]
// 		lastFu := sum[i][1]
// 		if v > 0 {
// 			sum[i+1] = [2]int{lastZheng + 1, lastFu}
// 		} else if v < 0 {
// 			sum[i+1] = [2]int{lastZheng, lastFu + 1}
// 		} else {
// 			sum[i+1] = [2]int{lastZheng, lastFu}
// 		}
// 	}

// 	for i := range nums {
// 		for j := i; j < len(nums) && nums[j] != 0; j++ {
// 			// zheng := sum[j+1][0] - sum[i][0]
// 			fu := sum[j+1][1] - sum[i][1]

// 			if fu%2 == 0 {
// 				answer = max(answer, j-i+1)
// 			}
// 		}
// 	}

// 	return answer
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{1, 2}
	fmt.Println(getMaxLen(nums))
}
