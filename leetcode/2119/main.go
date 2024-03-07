package main

import "fmt"

func isSameAfterReversals(num int) bool {
	r1 := reverse(num)
	return num == reverse(r1)
}

func reverse(num int) int {
	if num == 0 {
		return 0
	}

	answer := 0
	for num > 0 {
		answer = answer*10 + num%10
		num = num / 10
	}

	return answer
}

func main() {
	nums := 625
	fmt.Println(isSameAfterReversals(nums))
}
