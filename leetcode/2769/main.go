package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
	return num + t*2
}
func main() {
	num := 4
	t := 2
	fmt.Println(theMaximumAchievableX(num, t))
}
