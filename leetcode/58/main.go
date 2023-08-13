package main

import "fmt"

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	answer := 0
	for i >= 0 {
		if s[i] != ' ' {
			answer++
		} else {
			break
		}
		i--
	}

	return answer
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
