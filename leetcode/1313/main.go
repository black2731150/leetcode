package main

import "fmt"

func decompressRLElist(nums []int) []int {
	answer := []int{}
	for i := 0; i < len(nums); i += 2 {
		val := nums[i+1]
		n := nums[i]
		for j := 0; j < n; j++ {
			answer = append(answer, val)
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums))
}
