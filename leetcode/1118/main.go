package main

import "fmt"

func numberOfDays(year int, month int) int {
	answer := 0

	switch month {
	case 1:
		answer = 31
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			answer = 29
		} else {
			answer = 28
		}
	case 3:
		answer = 31
	case 4:
		answer = 30
	case 5:
		answer = 31
	case 6:
		answer = 30
	case 7:
		answer = 31
	case 8:
		answer = 31
	case 9:
		answer = 30
	case 10:
		answer = 31
	case 11:
		answer = 30
	case 12:
		answer = 31
	}

	return answer
}

func main() {
	year := 2000
	month := 2
	fmt.Println(numberOfDays(year, month))
}
