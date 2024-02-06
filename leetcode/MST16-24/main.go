package main

import "fmt"

func pairSums(nums []int, target int) [][]int {
	theMap := make(map[int]int)
	for _, v := range nums {
		theMap[v]++
	}
	answer := [][]int{}
	for _, v := range nums {
		s := v
		e := target - v

		if s == e {
			if num, find := theMap[e]; find {
				if num >= 2 {
					answer = append(answer, []int{e, s})
					theMap[e] -= 2
				}
			}
		} else {
			num1, find1 := theMap[e]
			num2, find2 := theMap[s]

			if find1 && find2 && num1 >= 1 && num2 >= 1 {
				answer = append(answer, []int{e, s})
				theMap[e]--
				theMap[s]--
			}
		}
	}

	return answer
}

func main() {
	nums := []int{4, 5, 6, 7, 8}
	target := 12
	fmt.Println(pairSums(nums, target))
}
