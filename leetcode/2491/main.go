package main

import (
	"fmt"
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	left, right := 0, len(skill)-1

	answer := int64(0)

	theNum := skill[left] + skill[right]
	for left < right {
		if skill[left]+skill[right] == theNum {
			answer += int64(skill[left] * skill[right])
			left++
			right--
		} else {
			return -1
		}
	}

	return int64(answer)
}
func main() {
	skill := []int{3, 2, 5, 1, 3, 4}
	fmt.Println(dividePlayers(skill))
}
