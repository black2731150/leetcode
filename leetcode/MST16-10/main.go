package main

import "fmt"

func maxAliveYear(birth []int, death []int) int {

	people := make([]int, 101)
	n := len(birth)
	m := 0
	answer := 0
	for i := 0; i < n; i++ {
		b := birth[i]
		d := death[i]

		for j := b; j <= d; j++ {
			people[j-1900]++
			if people[j-1900] >= m {
				if people[j-1900] == m {
					answer = min(answer, j)
				} else {
					answer = j
					m = people[j-1900]
				}
			}
		}
	}
	return answer
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	birth := []int{1972, 1908, 1915, 1957, 1960, 1948, 1912, 1903, 1949, 1977, 1900, 1957, 1934, 1929, 1913, 1902, 1903, 1901}
	death := []int{1997, 1932, 1963, 1997, 1983, 2000, 1926, 1962, 1955, 1997, 1998, 1989, 1992, 1975, 1940, 1903, 1983, 1969}
	fmt.Println(maxAliveYear(birth, death))
}
