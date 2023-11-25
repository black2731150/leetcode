package main

import "fmt"

func checkZeroOnes(s string) bool {
	n := len(s)
	numOfOne := 0
	numOfZero := 0
	maxOne, maxZero := 0, 0
	lastChar := s[0]
	if s[0] == '1' {
		numOfOne = 1
	} else {
		numOfZero = 1
	}

	for i := 1; i < n; i++ {
		if s[i] == lastChar {
			if lastChar == '1' {
				numOfOne++
			} else {
				numOfZero++
			}
		} else {
			maxOne = max(maxOne, numOfOne)
			maxZero = max(maxZero, numOfZero)
			if s[i] == '1' {
				lastChar = '1'
				numOfOne = 1
				numOfZero = 0
			} else {
				lastChar = '0'
				numOfZero = 1
				numOfOne = 0
			}
		}
	}
	maxOne = max(maxOne, numOfOne)
	maxZero = max(maxZero, numOfZero)

	return maxOne > maxZero
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	s := "1101"
	fmt.Println(checkZeroOnes(s))
}
