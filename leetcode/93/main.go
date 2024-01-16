package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	answer := []string{}

	tmp := []string{}
	var help func(i int)
	help = func(i int) {
		if i > len(s) {
			return
		}

		if i == len(s) {
			if len(tmp) == 4 {
				answer = append(answer, strings.Join(tmp, "."))
				return
			}
		}

		for j := i; j < i+3 && j < len(s); j++ {
			if is(s[i : j+1]) {
				tmp = append(tmp, s[i:j+1])
				help(j + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	help(0)
	return answer
}

func is(s string) bool {
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return num >= 0 && num <= 255
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
