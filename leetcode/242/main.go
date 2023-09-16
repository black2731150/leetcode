package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	theMap := make(map[byte]int, 26)

	for i := 'a'; i <= 'z'; i++ {
		theMap[byte(i)] = 0
	}

	for i := range s {
		theMap[s[i]]++
	}

	for i := range t {
		if _, find := theMap[t[i]]; find {
			theMap[t[i]]--
			if theMap[t[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "rat"
	t := "cat"
	fmt.Println(isAnagram(s, t))
}
