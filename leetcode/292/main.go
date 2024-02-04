package main

import "fmt"

func CanWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(CanWinNim(5))
}
