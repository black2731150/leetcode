package main

import "fmt"

func maxOperations(nums []int, k int) int {
	theMap := make(map[int]int)
	for _, v := range nums {
		theMap[v]++
	}

	answer := 0

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if k-nums[i] == nums[i] {
			if num, find := theMap[v]; find && num >= 2 {
				answer++
				theMap[v] -= 2
			}
		} else {
			firstNum, find1 := theMap[v]
			secondNum, find2 := theMap[k-v]
			if !(find1 && find2) {
				continue
			}

			if firstNum > 0 && secondNum > 0 {
				answer++
				theMap[v]--
				theMap[k-v]--
				i--
			}
		}
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	k := 5
	fmt.Println(maxOperations(nums, k))
}
