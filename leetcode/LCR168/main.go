package main

import "fmt"

var uglyNums = []int{1}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	index2, index3, index5 := 0, 0, 0

	for len(uglyNums) < 1690 {

		next2, next3, next5 := uglyNums[index2]*2, uglyNums[index3]*3, uglyNums[index5]*5
		nextUgly := min(next2, min(next3, next5))

		// 更新对应的指针
		if nextUgly == next2 {
			index2++
		}
		if nextUgly == next3 {
			index3++
		}
		if nextUgly == next5 {
			index5++
		}

		uglyNums = append(uglyNums, nextUgly)
	}
}

func nthUglyNumber(n int) int {
	return uglyNums[n-1]
}

func main() {
	n := 888
	fmt.Println(nthUglyNumber(n))
}
