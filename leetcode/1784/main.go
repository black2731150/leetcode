package main

import "fmt"

func checkOnesSegment(s string) bool {
	numOf1 := 0

	lianxu := false
	for i := 0; i < len(s); i++ {

		if numOf1 > 1 {
			return false
		}

		if s[i] == '1' {
			if lianxu {
				continue
			}
			numOf1++
			lianxu = true
		} else {
			lianxu = false
		}
	}
	return numOf1 <= 1
}

func main() {
	s := "110"
	fmt.Println(checkOnesSegment(s))
}
