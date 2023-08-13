package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	Len := len(s)

	S2T := make(map[byte]byte, Len)
	T2S := make(map[byte]byte, Len)

	for i := 0; i < Len; i++ {
		if S2T[s[i]] == 0 && T2S[t[i]] == 0 {
			S2T[s[i]] = t[i]
			T2S[t[i]] = s[i]
		} else {
			if S2T[s[i]] == t[i] && T2S[t[i]] == s[i] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
