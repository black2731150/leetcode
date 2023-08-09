package main

import "fmt"

func reformatNumber(number string) string {
	letters := []byte{}
	for i := range number {
		if number[i] == ' ' || number[i] == '-' {
			continue
		}
		letters = append(letters, byte(number[i]))
	}

	shengyu := len(letters) % 3
	answer := ""
	for i := 0; i+4 < len(letters); i = i + 3 {
		answer = answer + string(letters[i:i+3]) + "-"
	}
	if shengyu == 0 {
		answer = answer + string(letters[len(letters)-3]) + string(letters[len(letters)-2]) + string(letters[len(letters)-1])
	}

	if shengyu == 1 {
		answer = answer + string(letters[len(letters)-4]) + string(letters[len(letters)-3]) + "-" + string(letters[len(letters)-2]) + string(letters[len(letters)-1])
	}

	if shengyu == 2 {
		answer = answer + string(letters[len(letters)-2]) + string(letters[len(letters)-1])
	}

	return answer
}

func main() {
	number := "123 4-5678"
	fmt.Println(reformatNumber(number))
}
