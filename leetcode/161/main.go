package main

import "fmt"

func isOneEditDistance(s string, t string) bool {
	LenS, LenT := len(s), len(t)
	if LenS < LenT {
		return isOneEditDistance(t, s)
	}
	if LenS-LenT > 1 {
		return false
	}

	haveDiff := false
	for i, chT := range t {
		if s[i] != byte(chT) {
			haveDiff = true
			if LenS == LenT {
				return s[i+1:] == t[i+1:]
			}
			return s[i+1:] == t[i:]
		}
	}
	return haveDiff || LenS-LenT == 1
}

func main() {
	s := "ab"
	t := "abc"
	fmt.Println(isOneEditDistance(s, t))
}
