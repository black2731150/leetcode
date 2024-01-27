package main

import "fmt"

func longestSubsequence(arr []int, difference int) int {
	n := len(arr)
	theMap := make(map[int]int)

	max := 0
	for i := 0; i < n; i++ {
		theMap[arr[i]] = theMap[arr[i]-difference] + 1
		if theMap[arr[i]] > max {
			max = theMap[arr[i]]
		}
	}
	return max
}

func main() {
	arr := []int{1, 2, 3, 4}
	difference := 1
	fmt.Println(longestSubsequence(arr, difference))
}
