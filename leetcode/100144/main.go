package main

func findPeaks(mountain []int) []int {

	answer := []int{}
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			answer = append(answer, i)
		}
	}
	return answer
}
