package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	answer := 0
	for _, v := range hours {
		if v >= target {
			answer++
		}
	}

	return answer
}

func main() {
	hours := []int{0, 1, 2, 3, 4, 5}
	target := 2
	fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))
}
