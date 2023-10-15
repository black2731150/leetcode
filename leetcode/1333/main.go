package main

import (
	"fmt"
	"sort"
)

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {

	answer := [][]int{}

	for _, restaurant := range restaurants {
		if restaurant[3] > maxPrice {
			continue
		}

		if restaurant[4] > maxDistance {
			continue
		}

		if veganFriendly == 1 {
			if restaurant[2] == 0 {
				continue
			}
		}

		answer = append(answer, restaurant)
	}

	sort.Slice(answer, func(i, j int) bool {

		if answer[i][1] == answer[j][1] {
			return answer[i][0] > answer[j][0]
		}

		return answer[i][1] > answer[j][1]

	})

	trueAnswer := []int{}
	for i := range answer {
		trueAnswer = append(trueAnswer, answer[i][0])
	}

	return trueAnswer
}

func main() {
	resturants := [][]int{
		{1, 2, 0, 4, 5},
		{2, 3, 1, 5, 6},
		{3, 4, 1, 7, 8},
		{4, 1, 0, 3, 4},
	}

	veganFriendly := 1

	maxPrice := 7

	maxDistance := 7

	fmt.Println(filterRestaurants(resturants, veganFriendly, maxPrice, maxDistance))
}
