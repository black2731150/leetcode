package main

import "fmt"

func permute(nums []int) [][]int {
	answer := [][]int{}

	generateNums(nums, []int{}, &answer)
	return answer
}

func generateNums(nums []int, dangqian []int, answer *[][]int) {
	if len(nums) == 0 {
		*answer = append(*answer, dangqian)
		return
	}

	for i, num := range nums {

		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

		generateNums(newNums, append(dangqian, num), answer)
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
