package main

import "fmt"

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	n := len(calories)
	answer := 0
	for i := 0; i < n-k+1; i++ {
		t := 0
		for j := i; j < i+k; j++ {
			t += calories[j]
		}

		if t > upper {
			answer++
		}

		if t < lower {
			answer--
		}
	}

	return answer
}

func main() {
	calories := []int{1, 2, 3, 4, 5}
	k := 1
	lower := 3
	upper := 3
	fmt.Println(dietPlanPerformance(calories, k, lower, upper))
}
