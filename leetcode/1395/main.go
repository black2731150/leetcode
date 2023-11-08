package main

import "fmt"

func numTeams(rating []int) int {
	n := len(rating)
	answer := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if rating[i] < rating[j] {
				for k := j + 1; k < n; k++ {
					if rating[j] < rating[k] {
						answer++
					}
				}
			}

			if rating[i] > rating[j] {
				for k := j + 1; k < n; k++ {
					if rating[j] > rating[k] {
						answer++
					}
				}
			}

		}
	}
	return answer
}

func main() {
	rating := []int{1, 23, 4, 5, 6}
	fmt.Println(numTeams(rating))
}
