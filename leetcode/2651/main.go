package main

import "fmt"

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

func main() {
	arrivanTima := 4
	delayedTime := 5
	fmt.Println(findDelayedArrivalTime(arrivanTima, delayedTime))
}
