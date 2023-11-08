package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	theHashMap := make(map[string]int, n-1)

	answer := []string{}

	for i := 0; i <= n-10; i++ {
		theHashMap[s[i:i+10]]++
	}

	for k, v := range theHashMap {
		if v > 1 {
			answer = append(answer, k)
		}
	}

	return answer
}

func main() {
	s := "AAAAAAAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))

}
