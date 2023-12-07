package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}

	sMap, tMap := make(map[byte]int), make(map[byte]int)
	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for k := range sMap {
		if sMap[k] != tMap[k] {
			return false
		}
	}

	return true
}

func main() {
	s := "a"
	t := "t"
	fmt.Println(isAnagram(s, t))
}
