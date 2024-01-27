package main

import "fmt"

func maxRepOpt1(text string) int {

	return max(help(text), help(reverseString(text)))
}

func help(text string) int {
	letterNums := make([]int, 26)
	for i := range text {
		letterNums[text[i]-'a']++
	}

	answer := 0
	for i := 0; i < len(text); i++ {
		ch := text[i]
		num := letterNums[text[i]-'a']
		num--
		once := true
		have := 1
		for j := i + 1; j < len(text); j++ {
			if text[j] == ch {
				if num > 0 {
					num--
					have++
				}
			} else {
				if num > 0 && once {
					num--
					have++
					once = false
				} else {
					break
				}
			}
		}
		answer = max(answer, have)
	}

	return answer
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	test := "aaabbaaa"
	fmt.Println(maxRepOpt1(test))
}
