package main

import (
	"fmt"
	"strconv"
)

func maximum69Number(num int) int {
	numstr := []byte(strconv.Itoa(num))
	for i := range numstr {
		if numstr[i] == '6' {
			numstr[i] = '9'
			break
		}
	}
	answer, _ := strconv.Atoi(string(numstr))
	return answer
}

func main() {
	num := 9996
	fmt.Println(maximum69Number(num))
}
