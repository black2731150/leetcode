package main

import "fmt"

func numDecodings(s string) int {
<<<<<<< HEAD

	if len(s) == 0 {
=======
	if len(s) == 0 || s[0] == '0' {
>>>>>>> 3ce76188343f97bc79e7a13e53ed21b8afcfbd28
		return 0
	}

	if len(s) == 1 {
<<<<<<< HEAD
		if s[0] != '0' {
=======
		if s[0] != 0 {
>>>>>>> 3ce76188343f97bc79e7a13e53ed21b8afcfbd28
			return 1
		} else {
			return 0
		}
	}

	bp := make([]int, len(s))
<<<<<<< HEAD
	if s[0] != '0' {
		bp[0] = 1
	}

	if help(s[0], s[1]) {
		if s[1] != '0' {
			bp[1] = 2
		} else {
			bp[1] = 1
		}
	} else {
		if s[0] == '0' {
			bp[1] = 0
		}

		if s[0] == '2' && s[1] >= '7' {
			bp[1] = 1
		}

		if s[0] >= '3' {
			bp[1] = 1
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '0' {
				return 0
			}
			bp[i] = bp[i-2]
		} else {
			if help(s[i-1], s[i]) {
				if s[i-1] != '0' && s[i-1] <= '2' {
					bp[i] = bp[i-2] + bp[i-1]
				} else {
					bp[i] = bp[i-1]
				}
			} else {

			}
		}
	}

	fmt.Println(bp)
	return bp[len(bp)-1]
}

func help(ch1, ch2 byte) bool {
	// 第一个字符不能是0
	if ch1 == '0' {
		return false
	}

	// 组合不能超过26
	if ch1 == '2' && ch2 >= '7' {
		return false
	}

	// 组合不能超过30
	if ch1 >= '3' {
		return false
	}

	return true
=======
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
>>>>>>> 3ce76188343f97bc79e7a13e53ed21b8afcfbd28
}
