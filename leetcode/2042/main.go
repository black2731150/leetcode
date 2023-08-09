package main

import (
	"fmt"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	strings := strings.Split(s, " ")
	nums := []int{}
	for _, v := range strings {
		num, err := strconv.Atoi(v)
		if err == nil {
			nums = append(nums, num)
		}
	}

	for i := 0; i+1 < len(nums); i++ {
		if nums[i] < nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
	fmt.Println(areNumbersAscending(s))
}
