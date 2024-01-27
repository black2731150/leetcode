package main

import "fmt"

func findDuplicate(nums []int) int {
	theMap := make(map[int]struct{})
	for _, v := range nums {
		if _, find := theMap[v]; find {
			return v
		} else {
			theMap[v] = struct{}{}
		}
	}
	return 0
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
