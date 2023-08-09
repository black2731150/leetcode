package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	answer := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i]+duration < timeSeries[i+1] {
			answer = answer + duration
		} else {
			answer = answer + timeSeries[i+1] - timeSeries[i]
		}
	}

	answer = answer + duration
	return answer
}
