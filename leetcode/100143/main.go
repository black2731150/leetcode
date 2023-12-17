package main

import "fmt"

func countTestedDevices(batteryPercentages []int) int {

	answer := 0
	x := 0

	for i := range batteryPercentages {

		batteryPercentages[i] = max(0, batteryPercentages[i]-x)

		if batteryPercentages[i] == 0 {
			continue
		}

		answer++
		x++
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	batteryPercentages := []int{0, 1, 2}
	fmt.Println(countTestedDevices(batteryPercentages))
}
