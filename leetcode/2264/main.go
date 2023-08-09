package main

import "fmt"

func largestGoodInteger(num string) string {
	answer := byte('0' - 1)

	for i := 0; i < len(num)-2; i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if answer < num[i] {
				answer = num[i]
			}
		}
	}
	if answer == byte('0'-1) {
		return ""
	}
	return string([]byte{answer, answer, answer})
}

func main() {
	num := "222"
	fmt.Println(largestGoodInteger(num))
}
