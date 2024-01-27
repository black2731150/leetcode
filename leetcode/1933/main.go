package main

import "fmt"

func isDecomposable(s string) bool {
	now := s[0]
	num := 1
	hasTow := false

	for i := 1; i < len(s); i++ {
		if s[i] == now {
			num++
			if num == 3 {
				num = 0
			}
		} else {
			if num == 0 {
				now = s[i]
				num = 1
			} else if num == 2 {
				if !hasTow {
					now = s[i]
					num = 1
					hasTow = true
				} else {
					return false
				}
			} else if num == 1 {
				return false
			}
		}
	}
	if num == 1 {
		return false
	}
	if num == 2 {
		hasTow = true
	}

	return true && hasTow
}

func main() {
	s := "000111000"
	fmt.Println(isDecomposable(s))
}
