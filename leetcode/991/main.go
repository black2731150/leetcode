package main

import "fmt"

func brokenCalc(startValue int, target int) int {
	answer := 0
	for target > startValue {
		answer++
		if target%2 == 1 {
			target++
		} else {
			target /= 2
		}
	}

	return answer + (startValue - target)
}

func main() {
	startValue := 2
	taget := 3
	fmt.Println(brokenCalc(startValue, taget))
}
