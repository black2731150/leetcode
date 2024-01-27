package main

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
	chars := make([]int, 26)
	for i := range s {
		chars[s[i]-'a']++
	}

	answer := []byte{}
	lianxu := 0
	for i := len(chars) - 1; i >= 0; {
		if chars[i] > 0 && lianxu < repeatLimit {
			answer = append(answer, byte(i+'a'))
			chars[i]--
			lianxu++
		} else {
			if chars[i] == 0 {
				lianxu = 0
				i--
			} else {
				find := false
				for j := i - 1; j >= 0; j-- {
					if chars[j] > 0 {
						chars[j]--
						answer = append(answer, byte(j+'a'))
						lianxu = 0
						find = true
						break
					}
				}

				if !find {
					break
				}
			}
		}
	}

	return string(answer)
}

func main() {
	s := "cczazcc"
	repeatLimit := 3
	fmt.Println(repeatLimitedString(s, repeatLimit))
}
