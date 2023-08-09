package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {

	theMap := map[int]int{}
	for i := range score {
		theMap[score[i]] = i
	}

	var copyScore = make([]int, len(score))
	copy(copyScore, score)

	sort.Ints(copyScore)

	answer := make([]string, len(score))
	for i := range copyScore {
		if i >= len(copyScore)-3 {
			if i == len(copyScore)-1 {
				answer[theMap[copyScore[i]]] = "Gold Medal"
			}
			if i == len(copyScore)-2 {
				answer[theMap[copyScore[i]]] = "Silver Medal"
			}
			if i == len(copyScore)-3 {
				answer[theMap[copyScore[i]]] = "Bronze Medal"
			}
		} else {
			answer[theMap[copyScore[i]]] = strconv.Itoa(len(copyScore) - i)
		}
	}
	return answer
}

func main() {
	score := []int{5, 4}
	fmt.Println(findRelativeRanks(score))
}
