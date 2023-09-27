package main

import "fmt"

func passThePillow(n int, time int) int {

	yiCiXunHuan := 2 * (n - 1)
	if yiCiXunHuan < time {
		time = time % yiCiXunHuan
	}

	if time < n {
		return time + 1
	} else {
		return 2*(n-1) + 1 - time
	}
}

func main() {
	n := 5
	time := 9
	fmt.Println(passThePillow(n, time))
}
