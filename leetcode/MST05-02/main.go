package main

import "fmt"

func printBin(num float64) string {
	answer := []byte{'0', '.'}

	for i := 0; i < 30; i++ {

		if num == 0 {
			return string(answer)
		}

		num = num * 2
		if num >= 1 {
			num = num - 1
			answer = append(answer, '1')
		} else {
			answer = append(answer, '0')
		}
	}

	return "ERROR"
}

func main() {
	num := 0.625
	fmt.Println(printBin(num))
}
