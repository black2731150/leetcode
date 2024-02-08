package main

import "fmt"

func calculateTime(keyboard string, word string) int {
	theMap := make(map[byte]int)
	for i := 0; i < len(keyboard); i++ {
		theMap[keyboard[i]] = i
	}

	answer := 0
	index := 0
	for i := 0; i < len(word); i++ {
		answer += abs(theMap[word[i]] - index)
		index = theMap[word[i]]
	}

	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	keyboard := "abcdefghijklmnopqrstuvwxyz"
	word := "cba"
	fmt.Println(calculateTime(keyboard, word))
}
