package main

import (
	"fmt"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"
	start, _ := time.Parse(layout, date1)
	end, _ := time.Parse(layout, date2)

	duration := end.Sub(start)
	answer := int(duration.Seconds()) / (86400)

	if answer < 0 {
		return -answer
	}
	return answer
}

func main() {
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	fmt.Println(daysBetweenDates(date1, date2))
}
