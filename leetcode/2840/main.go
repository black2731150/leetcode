package main

import "fmt"

func checkStrings(s1 string, s2 string) bool {
	s2JiMap := make(map[byte]int)
	s2OuMap := make(map[byte]int)

	for i := 0; i < len(s2); i++ {
		if i%2 == 0 {
			s2OuMap[s1[i]]--
			s2OuMap[s2[i]]++
		} else {
			s2JiMap[s1[i]]--
			s2JiMap[s2[i]]++
		}
	}

	for _, v := range s2JiMap {
		if v != 0 {
			return false
		}
	}

	for _, v := range s2OuMap {
		if v != 0 {
			return false
		}
	}

	return true

}

func main() {
	s1 := "abcdba"
	s2 := "cabdab"
	fmt.Println(checkStrings(s1, s2))
}
