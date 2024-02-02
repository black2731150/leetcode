package main

import "fmt"

func replaceSpaces(S string, length int) string {
	answer := []byte{}
	for i := 0; i < length; i++ {
		if S[i] != ' ' {
			answer = append(answer, S[i])
		} else {
			answer = append(answer, []byte{'%', '2', '0'}...)
		}
	}

	return string(answer)
}

func main() {
	S := "Mr John Smith    "
	lenght := 13
	fmt.Println(replaceSpaces(S, lenght))
}
