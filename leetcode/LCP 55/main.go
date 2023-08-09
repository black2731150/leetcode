package main

func getMinimumTime(time []int, fruits [][]int, limit int) int {
	answer := 0
	for i := range fruits {
		cishu := 0
		if x := fruits[i][1] % limit; x == 0 {
			cishu = fruits[i][1] / limit
		} else {
			cishu = fruits[i][1]/limit + 1
		}
		answer = answer + time[fruits[i][0]]*cishu
	}

	return answer
}
