package main

import (
	"fmt"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	leaveMap := make(map[int]int)

	treapIndex := 0
	for i := 0; i <= 1000 && treapIndex < len(trips); i++ {

		if leave, find := leaveMap[i]; find {
			capacity += leave
			leaveMap[i] = 0
		}

		if trips[treapIndex][1] == i {
			upPeople := trips[treapIndex][0]
			to := trips[treapIndex][2]

			if capacity-upPeople >= 0 {
				capacity -= upPeople
				leaveMap[to] += upPeople
				treapIndex++
				i--
			} else {
				return false
			}

		}
	}

	return true
}

func main() {
	trips := [][]int{{4, 1, 3}, {5, 3, 7}}
	capacity := 6
	fmt.Println(carPooling(trips, capacity))
}
