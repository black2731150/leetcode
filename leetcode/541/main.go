package main

import "fmt"

func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(s); i = i + 2*k {
		if i+k < len(s) {
			for i2, j := 0, len(bytes[i:i+k])-1; i2 < j; i2, j = i2+1, j-1 {
				bytes[i : i+k][i2], bytes[i : i+k][j] = bytes[i : i+k][j], bytes[i : i+k][i2]
			}
		} else {
			for i2, j := 0, len(bytes[i:])-1; i2 < j; i2, j = i2+1, j-1 {
				bytes[i:][i2], bytes[i:][j] = bytes[i:][j], bytes[i:][i2]
			}
		}
	}

	return string(bytes)
}

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(s, k)
}
