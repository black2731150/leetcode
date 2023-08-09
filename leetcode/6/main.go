package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	mat := make([][]byte, numRows)
	t, x := 2*(numRows-1), 0

	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < numRows-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}
