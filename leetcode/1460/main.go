package main

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}

	sort.Ints(target)
	sort.Ints(arr)

	for i, v := range arr {
		if v != target[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := []int{5, 4, 3, 2, 1}
	fmt.Println(canBeEqual(target, arr))
}
