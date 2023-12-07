package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	index := 1
	answer := []int{}

	for _, v := range arr {
		if v == index {
			index++
		} else {
			for index != v {
				answer = append(answer, index)
				index++
			}
			index++
		}
	}

	if len(answer) < k {
		return arr[len(arr)-1] + k - len(answer)
	}
	return answer[k-1]
}

func main() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	fmt.Println(findKthPositive(arr, k))
}
