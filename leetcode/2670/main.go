package main

import "fmt"

func distinctDifferenceArray(nums []int) []int {
	answer := []int{}
	for i := 0; i < len(nums); i++ {
		x := diffNum(nums[:i+1]) - diffNum(nums[i+1:])
		answer = append(answer, x)
	}
	return answer
}

func diffNum(nums []int) int {
	answer := 0
	themap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := themap[v]; !ok {
			answer++
			themap[v] = true
		}
	}
	for k := range themap {
		delete(themap, k)
	}

	return answer
}

func main() {
	nums := []int{3, 2, 3, 4, 2}
	fmt.Println(distinctDifferenceArray(nums))

}
