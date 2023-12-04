package main

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)

	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] <= mass {
			mass += asteroids[i]
		} else {
			return false
		}
	}

	return true
}

func main() {
	mass := 10
	asteroids := []int{4, 5, 8, 11, 23}
	fmt.Println(asteroidsDestroyed(mass, asteroids))
}
