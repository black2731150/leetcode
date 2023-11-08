package main

import "fmt"

func numberOfBeams(bank []string) int {
	zheYiCeng := 0
	shangYiCeng := 0

	answer := 0
	for _, v := range bank {
		for _, ch := range v {
			if ch == '1' {
				zheYiCeng++
			}
		}
		if zheYiCeng == 0 {
			continue
		} else {
			answer = answer + shangYiCeng*zheYiCeng
			shangYiCeng = zheYiCeng
			zheYiCeng = 0
		}
	}
	return answer
}

func main() {
	bank := []string{"011001", "000000", "010100", "001000"}
	fmt.Println(numberOfBeams(bank))
}
