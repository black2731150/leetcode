package main

import "fmt"

func maximum(a int, b int) int {
	answer := a - b
	k := int(answer>>63) & 1
	fk := k ^ 1

	return a*fk + b*k
}

func main() {
	a := 32242423
	b := 2324242
	fmt.Println(maximum(a, b))
}
