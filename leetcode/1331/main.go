package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	var sortedArr []int = make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	theSortedMap := map[int]int{}
	index := 1
	for _, v := range sortedArr {
		if _, ok := theSortedMap[v]; !ok {
			theSortedMap[v] = index
			index++
		}
	}

	answer := []int{}
	for i := 0; i < len(arr); i++ {
		answer = append(answer, theSortedMap[arr[i]])
	}

	return answer

}

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr))
}
