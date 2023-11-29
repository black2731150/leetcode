package main

import (
	"fmt"
	"sort"
)

func maxSumDivThree(nums []int) int {

	answer := 0

	num2 := []int{}
	num1 := []int{}
	for _, v := range nums {
		if v%3 == 0 {
			answer += v
			continue
		}
		if v%3 == 2 {
			num2 = append(num2, v)
			continue
		}
		if v%3 == 1 {
			num1 = append(num1, v)
			continue
		}
	}

	sort.Ints(num1)
	sort.Ints(num2)

	for len(num1) >= 3 && len(num2) >= 3 {
		one := num1[len(num1)-1] + num2[len(num2)-1]
		tow := num1[len(num1)-1] + num1[len(num1)-2] + num1[len(num1)-3]
		three := num2[len(num2)-1] + num2[len(num2)-2] + num2[len(num2)-3]

		if one > tow {
			if one > three {
				answer += one
				num1 = num1[:len(num1)-1]
				num2 = num2[:len(num2)-1]
			} else {
				answer += three
				num2 = num2[:len(num2)-3]
			}
		} else {
			if tow > three {
				answer += tow
				num1 = num1[:len(num1)-3]
			} else {
				answer += three
				num2 = num2[:len(num2)-3]
			}
		}
	}

	for (len(num1) >= 0 && len(num1) < 3) || (len(num2) >= 0 && len(num2) < 3) {
		if len(num1) < 3 && len(num2) < 3 {
			break
		}

		if len(num1) == 0 || len(num2) == 0 {
			break
		}

		if len(num1) >= 3 {
			one := num1[len(num1)-1] + num2[len(num2)-1]
			tow := num1[len(num1)-1] + num1[len(num1)-2] + num1[len(num1)-3]
			if one > tow {
				answer += one
				num1 = num1[:len(num1)-1]
				num2 = num2[:len(num2)-1]
			} else {
				answer += tow
				num1 = num1[:len(num1)-3]
			}
		}

		if len(num2) >= 3 {
			one := num1[len(num1)-1] + num2[len(num2)-1]
			tow := num2[len(num2)-1] + num2[len(num2)-2] + num2[len(num2)-3]

			if one > tow {
				answer += one
				num1 = num1[:len(num1)-1]
				num2 = num2[:len(num2)-1]
			} else {
				answer += tow
				num2 = num2[:len(num2)-3]
			}
		}
	}

	for len(num1) > 0 && len(num2) > 0 {
		answer += num1[len(num1)-1] + num2[len(num2)-1]
		num1 = num1[:len(num1)-1]
		num2 = num2[:len(num2)-1]
	}

	for len(num1) >= 3 {
		answer += num1[len(num1)-1] + num1[len(num1)-2] + num1[len(num1)-3]
		num1 = num1[:len(num1)-3]
	}

	for len(num2) >= 3 {
		answer += num2[len(num2)-1] + num2[len(num2)-2] + num2[len(num2)-3]
		num2 = num2[:len(num2)-3]
	}

	return answer
}

func main() {
	nums := []int{5, 2, 2, 2}
	fmt.Println(maxSumDivThree(nums))
}
