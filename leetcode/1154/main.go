package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	times := strings.Split(date, "-")
	year, _ := strconv.Atoi(times[0])
	month, _ := strconv.Atoi(times[1])
	day, _ := strconv.Atoi(times[2])

	theNums := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		theNums[1] = 29
	}

	answer := 0
	for i := 0; i < month-1; i++ {
		answer += theNums[i]
	}
	answer += day
	return answer
}

func main() {
	date := "2019-01-27"
	fmt.Println(dayOfYear(date))
}
