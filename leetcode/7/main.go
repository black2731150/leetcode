package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	fuhao := true
	if x < 0 {
		fuhao = false
		x = -x
	}
	xstr := strconv.Itoa(x)
	newStr := []byte{}
	for i := len(xstr) - 1; i >= 0; i-- {
		newStr = append(newStr, xstr[i])
	}
	num, _ := strconv.Atoi(string(newStr))

	if num > math.MaxInt32 {
		return 0
	}

	if fuhao {
		return num
	} else {
		return -num
	}
}

func main() {
	x := 1534236469
	fmt.Println(reverse(x))
}
