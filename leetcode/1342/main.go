package main

func numberOfSteps(num int) int {
	answer := 0
	for i := num; i > 0; i = i / 2 {
		if i%2 == 1 {
			answer++
			i = i - 1
		}
		if i != 0 {
			answer++
		}

	}
	return answer
}
