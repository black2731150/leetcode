package main

import (
	"fmt"
	"sort"
)

func removeDigit(number string, digit byte) string {
	answers := []string{}

	for i := range number {
		if number[i] == digit {
			answers = append(answers, number[:i]+number[i+1:])
		}
	}

	sort.Strings(answers)
	return answers[len(answers)-1]
}

func main() {
	number := "123"
	digit := '3'
	fmt.Println(removeDigit(number, byte(digit)))
}
