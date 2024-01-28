package main

import "fmt"

func canBeEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	c1, c2, c3, c4 := s1[0], s1[1], s1[2], s1[3]

	if string([]byte{c3, c2, c1, c4}) == s2 {
		return true
	}

	if string([]byte{c1, c4, c3, c2}) == s2 {
		return true
	}

	if string([]byte{c3, c4, c1, c2}) == s2 {
		return true
	}

	return false

}

func main() {
	s1 := "abcd"
	s2 := "cbad"
	fmt.Println(canBeEqual(s1, s2))
}
