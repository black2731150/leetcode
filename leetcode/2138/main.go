package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	answer := []string{}
	zushu := len(s) / k
	shengyu := len(s) % k
	for i := 0; i < zushu; i++ {
		answer = append(answer, s[i*k:(i+1)*k])
	}

	if shengyu != 0 {
		last := s[zushu*k:]
		for i := 0; i < k-shengyu; i++ {
			last = last + string(fill)
		}
		answer = append(answer, last)
	}

	return answer
}

func main() {
	s := "abcdefghij"
	k := 3
	fill := 'x'
	fmt.Println(divideString(s, k, byte(fill)))
}
