package main

import (
	"fmt"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	targetMap := make(map[byte]int)
	answer := ""
	for i := range licensePlate {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			targetMap[licensePlate[i]]++
		}
	}

	for _, word := range words {
		tmp := make(map[byte]int)
		for j := range word {
			tmp[word[j]]++
		}

		success := true
		for k := range targetMap {
			if targetMap[k]-tmp[k] <= 0 {
				continue
			} else {
				success = false
				break
			}
		}

		if success {
			if len(answer) == 0 {
				answer = word
			}

			if len(word) < len(answer) {
				answer = word
			}
		}
	}

	return answer
}

func main() {
	licensePlate := "1s3 PSt"
	words := []string{"step", "steps", "stripe", "stepple"}
	fmt.Println(shortestCompletingWord(licensePlate, words))
}
