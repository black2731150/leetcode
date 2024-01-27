package main

import "fmt"

func maximumSwap(num int) int {
	arr := []int{}
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := 0; i < len(arr); i++ {
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] >= arr[maxIndex] && arr[j] > arr[i] {
				maxIndex = j
			}
		}

		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
			break
		}
	}

	answer := 0
	for _, v := range arr {
		answer = answer*10 + v
	}

	return answer
}

func main() {
	num := 2736
	fmt.Println(maximumSwap(num))
}
