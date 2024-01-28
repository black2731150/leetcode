package main

import "fmt"

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if jug1Capacity+jug2Capacity == targetCapacity || jug1Capacity == targetCapacity || jug2Capacity == targetCapacity {
		return true
	}

	if jug1Capacity == jug2Capacity || jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}

	return targetCapacity%gcd(jug1Capacity, jug2Capacity) == 0

}

func gcd(x, y int) int {
	if x < y {
		return gcd(y, x)
	}

	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func main() {
	j1 := 3
	j2 := 5
	t := 4
	fmt.Println(canMeasureWater(j1, j2, t))
}
