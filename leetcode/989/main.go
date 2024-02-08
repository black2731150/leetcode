package main

import "fmt"

func addToArrayForm(num []int, k int) []int {

	n := len(num) - 1
	count := 0
	for k > 0 {
		x := k % 10
		k = k / 10

		if n >= 0 {
			num[n] = num[n] + x + count
			count = 0
			if num[n] >= 10 {
				num[n] = num[n] - 10
				count = 1
			}
			n--
		} else {
			v := x + count
			count = 0
			if v >= 10 {
				v = v - 10
				count = 1
			}
			num = append([]int{v}, num...)
		}
	}

	for n >= 0 {
		num[n] = num[n] + count
		count = 0
		if num[n] >= 10 {
			num[n] = num[n] - 10
			count = 1
		}
		n--
	}

	if count == 1 {
		num = append([]int{1}, num...)
	}

	return num
}

func main() {
	num := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	k := 1
	fmt.Println(addToArrayForm(num, k))
}
