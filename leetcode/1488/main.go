package main

func avoidFlood(rains []int) []int {
	n := len(rains)

	answer := []int{}

	fullLake := make(map[int]int, n)

	for i, v := range rains {
		if v != 0 {
			answer = append(answer, -1)
			fullLake[v]++
		} else {

		}
	}

	return answer
}
