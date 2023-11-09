package main

import "fmt"

func averageValue(nums []int) int {
	numberOfOu := 0
	sum := 0
	for _, v := range nums {
		if v%2 == 0 {
			if v%3 == 0 {
				numberOfOu++
				sum += v
			}
		}
	}

	if numberOfOu == 0 {
		return 0
	}
	return sum / numberOfOu
}

func main() {
	nums := []int{4, 6, 7, 12, 18}
	fmt.Println(averageValue(nums))
}
