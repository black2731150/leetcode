package main

import (
	"fmt"
	"strconv"
)

func countNicePairs(nums []int) int {
	mod := 1e9 + 7
	n := len(nums)
	answer := 0
	revNums := make([]int, n)
	for i, v := range nums {
		revNums[i] = rev(v)
	}

	theMap := make(map[int]int)

	for i := range nums {
		theMap[nums[i]-revNums[i]]++
	}

	for _, v := range theMap {
		answer = (answer + v*(v-1)/2) % int(mod)
	}

	return answer
}

func rev(num int) int {
	strNum := strconv.Itoa(num)
	byteNum := []byte(strNum)
	for i, j := 0, len(byteNum)-1; i < j; i, j = i+1, j-1 {
		byteNum[i], byteNum[j] = byteNum[j], byteNum[i]
	}
	strNum = string(byteNum)
	answer, _ := strconv.Atoi(strNum)
	return answer
}

func main() {
	nums := []int{3, 4, 7, 8, 9, 1}
	fmt.Println(countNicePairs(nums))
}
