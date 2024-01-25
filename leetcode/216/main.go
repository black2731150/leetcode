package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tmp := []int{}
	answer := [][]int{}
	var help func(start int, t int)
	help = func(start int, t int) {
		if t < 0 {
			return
		}

		if t == 0 && len(tmp) == k {
			copyTmp := make([]int, len(tmp))
			copy(copyTmp, tmp)
			answer = append(answer, copyTmp)
		}

		for i := start; i < 9; i++ {
			if nums[i] > n {
				break
			}

			tmp = append(tmp, nums[i])
			help(i+1, t-nums[i])
			tmp = tmp[:len(tmp)-1]
		}
	}
	help(0, n)
	return answer
}

func main() {
	k := 3
	n := 9
	fmt.Println(combinationSum3(k, n))
}
