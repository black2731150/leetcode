package main

import "fmt"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	answer := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					answer++
				}
			}
		}
	}
	return answer
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	fmt.Println(countGoodTriplets(arr, a, b, c))
}
