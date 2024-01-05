package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return t.Weekday().String()
}

func main() {
	day := 31
	month := 8
	year := 2019
	fmt.Println(dayOfTheWeek(day, month, year))
}
