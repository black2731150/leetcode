package main

import "fmt"

func minimumSeconds(nums []int) int {

	theMap := make(map[int][]int)
	for i, v := range nums {
		theMap[v] = append(theMap[v], i)
	}
	n := len(nums)

	answer := int(1e6)

	for _, indexs := range theMap {
		m := 0
		if len(indexs) == 1 {
			m = max(m, n)
		} else {
			for i := 0; i < len(indexs)-1; i++ {
				m = max(m, indexs[i+1]-indexs[i])
			}
			m = max(m, indexs[0]-indexs[len(indexs)-1]+n)
		}

		answer = min(answer, m/2)
	}

	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 1, 3, 3, 2}
	fmt.Println(minimumSeconds(nums))
}
