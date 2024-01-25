package main

import "fmt"

func findFinalValue(nums []int, original int) int {
	theMap := make(map[int]struct{})
	for _, v := range nums {
		theMap[v] = struct{}{}
	}

	for {
		if _, find := theMap[original]; find {
			original = 2 * original
		} else {
			break
		}
	}

	return original
}

func main() {
	nums := []int{5, 3, 6, 1, 12}
	original := 3
	fmt.Println(findFinalValue(nums, original))
}
