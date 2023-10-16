package main

import (
	"fmt"
	"strings"
)

func sumDistance(nums []int, s string, d int) int {
	Strings := strings.Split(s, "")

	for ; d > 0; d-- {
		for i := range nums {
			if Strings[i] == "L" {
				nums[i]--
			} else {
				nums[i]++
			}

			if i > 0 {
				if nums[i] == nums[i-1] {
					if Strings[i] == "L" {
						Strings[i] = "R"
					} else {
						Strings[i] = "L"
					}

					if Strings[i-1] == "L" {
						Strings[i-1] = "R"
					} else {
						Strings[i-1] = "L"
					}
				}
			}
		}
	}

	answer := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			answer = (answer + abs(nums[i], nums[j])) % 1000000007
		}
	}

	return answer
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}

func main() {
	nums := []int{-2, 0, 2}
	s := "RLL"
	d := 3
	fmt.Println(sumDistance(nums, s, d))
}
