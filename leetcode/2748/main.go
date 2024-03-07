package main

import "fmt"

func countBeautifulPairs(nums []int) int {
	answer := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if gcd(getFristNum(nums[i]), getLastnNum(nums[j])) == 1 {
				answer++
			}
		}
	}

	return answer
}

func getLastnNum(x int) int {
	return x % 10
}

func getFristNum(x int) int {
	for x != 0 && x/10 != 0 {
		x = x / 10
	}

	return x
}

func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}

	return y
}

func main() {
	nums := []int{31, 25, 72, 79, 74}
	fmt.Println(countBeautifulPairs(nums))
}
