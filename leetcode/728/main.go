package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	answer := []int{}
	for i := left; i <= right; i++ {
		if isTheRightNum(i) {
			answer = append(answer, i)
		}
	}
	return answer
}

func isTheRightNum(num int) bool {
	thenum := num
	for ; thenum > 0; thenum = thenum / 10 {
		x := thenum % 10
		if x == 0 {
			return false
		}

		if num%x == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	left := 47
	right := 85
	fmt.Println(selfDividingNumbers(left, right))
}
