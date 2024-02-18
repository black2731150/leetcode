package main

import (
	"fmt"
	"sort"
)

func longestDiverseString(a int, b int, c int) string {
	answer := []byte{}
	type point struct {
		letter byte
		num    int
	}

	arr := []point{{'a', a}, {'b', b}, {'c', c}}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num > arr[j].num
	})
	answer = append(answer, arr[0].letter)
	arr[0].num--
	if arr[0].num > 0 {
		answer = append(answer, arr[0].letter)
		arr[0].num--
	}

	for {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].num > arr[j].num
		})

		if answer[len(answer)-1] == arr[0].letter && answer[len(answer)-2] == arr[0].letter {
			if arr[1].num > 0 {
				answer = append(answer, arr[1].letter)
				arr[1].num--
			} else {
				break
			}
		} else {
			if arr[0].num >= 2 {
				answer = append(answer, arr[0].letter)
				answer = append(answer, arr[0].letter)
				arr[0].num -= 2
			} else if arr[0].num >= 1 {
				answer = append(answer, arr[0].letter)
				arr[0].num -= 1

			} else {
				break
			}
		}
	}

	return string(answer)
}

func main() {
	a, b, c := 1, 1, 7
	fmt.Println(longestDiverseString(a, b, c))
}
