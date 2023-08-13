package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	theMap := make(map[byte]int, 26)

	for i := 'a'; i <= 'z'; i++ {
		theMap[byte(i)] = 0
	}

	for i := range magazine {
		theMap[magazine[i]]++
	}

	for i := range ransomNote {
		theMap[ransomNote[i]]--
		if theMap[ransomNote[i]] < 0 {
			return false
		}
	}

	return true
}

func main() {
	ransoNote := "aa"
	magezine := "ab"
	fmt.Println(canConstruct(ransoNote, magezine))
}
