package main

import "fmt"

func longestConsecutive(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	theMap := make(map[int]bool)
	for i := range nums {
		theMap[nums[i]] = true
	}

	answer := 0
	tmp := 0

	for i := range theMap {

		if !theMap[i-1] {
			for j := i; theMap[j]; j++ {
				tmp++
			}
			answer = max(answer, tmp)
			tmp = 0
		}

	}

	answer = max(answer, tmp)

	return answer

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
