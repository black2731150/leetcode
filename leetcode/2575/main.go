package main

import "fmt"

func divisibilityArray(word string, m int) []int {
	num := 0
	answer := []int{}
	for i := 0; i < len(word); i++ {
		num = num*10 + int(word[i]-'0')
		if num%m == 0 {
			answer = append(answer, 1)
		} else {
			answer = append(answer, 0)
		}
		num = num % m
	}

	return answer
}

func main() {
	word := "998244353"
	m := 3
	fmt.Println(divisibilityArray(word, m))
}
