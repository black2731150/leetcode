package main

import "fmt"

func calculateTax(brackets [][]int, income int) float64 {
	var answer, x, all float64 = 0, 0, 0

	for _, v := range brackets {
		upper := v[0]
		percent := v[1]
		if income > upper {
			x = float64(upper) - all
			all += x
			answer += x * (float64(percent)) / 100
		} else {
			x = float64(income) - all
			answer += x * (float64(percent)) / 100
			break
		}
	}
	return answer
}

func main() {
	brackets := [][]int{{2, 7}, {3, 17}, {4, 37}, {7, 6}, {9, 83}, {16, 67}, {19, 29}}
	income := 18
	fmt.Println(calculateTax(brackets, income))
}
