package main

import "fmt"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	for i := range s1 {
		if tmp := s1[i:] + s1[:i]; tmp == s2 {
			return true
		}
	}

	return false
}

func main() {
	s1 := "waterbottle"
	s2 := "waterbottle"
	fmt.Println(isFlipedString(s1, s2))
}
