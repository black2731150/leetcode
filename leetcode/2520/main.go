package main

import "fmt"

func countDigits(num int) int {
	copyNum := num

	answer := 0
	for copyNum != 0 {
		x := copyNum % 10
		copyNum = copyNum / 10

		if num%x == 0 {
			answer++
		}
	}

	return answer
}

func main() {
	num := 1246
	fmt.Println(countDigits(num))
}
