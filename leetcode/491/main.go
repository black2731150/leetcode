package main

import (
	"fmt"
	"strconv"
)

func findSubsequences(nums []int) [][]int {
	answer := [][]int{}

	theMap := make(map[string]struct{})
	tmp := []int{}
	var help func(index int)
	help = func(index int) {
		if len(tmp) >= 2 {
			x := sliceToString(tmp)
			if _, find := theMap[x]; !find {
				theMap[x] = struct{}{}
				tmpCopy := make([]int, len(tmp))
				copy(tmpCopy, tmp)
				answer = append(answer, tmpCopy)
			}
		}

		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			if nums[i] >= tmp[len(tmp)-1] {
				tmp = append(tmp, nums[i])
				help(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		help(i + 1)
		tmp = tmp[:len(tmp)-1]
	}

	return answer
}

func sliceToString(slice []int) string {
	answer := ""
	for _, v := range slice {
		answer += strconv.Itoa(v) + ","
	}
	return answer
}

func main() {
	nums := []int{4, 5, 7, 7}
	fmt.Println(findSubsequences(nums))
}
