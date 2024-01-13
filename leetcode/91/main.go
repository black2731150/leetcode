package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		if s[0] != 0 {
			return 1
		} else {
			return 0
		}
	}

	bp := make([]int, len(s))
	bp[0] = 1

	if s[1] != '0' {
		bp[1] = bp[0]
	}
	if help(s[0], s[1]) {
		bp[1]++
	}

	for i := 2; i < len(s); i++ {
		if s[i] != '0' {
			bp[i] = bp[i-1]
		}

		if help(s[i-1], s[i]) {
			bp[i] += bp[i-2]
		} else if s[i] == '0' {
			return 0 // 当前字符是 '0' 且不能与前一个字符组合，没有解码方式
		}
	}

	return bp[len(s)-1]
}

func help(ch1, ch2 byte) bool {
	return !(ch1 == '0' || (ch1 == '2' && ch2 > '6') || ch1 > '2')
}

func main() {
	s := "12"
	fmt.Println(numDecodings(s))
}
