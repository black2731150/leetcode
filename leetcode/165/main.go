package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	v1Index, v2Index := 0, 0

	for v1Index < len(v1) || v2Index < len(v2) {
		v1Ch, v2Ch := "", ""
		if v1Index >= len(v1) {
			v1Ch = "0"
		} else {
			v1Ch = v1[v2Index]
		}

		if v2Index >= len(v2) {
			v2Ch = "0"
		} else {
			v2Ch = v2[v2Index]
		}

		v1Num, _ := strconv.Atoi(v1Ch)
		v2Num, _ := strconv.Atoi(v2Ch)

		if v1Num < v2Num {
			return -1
		}

		if v1Num > v2Num {
			return 1
		}

		v1Index++
		v2Index++
	}
	return 0
}

func main() {
	version1 := "1.01"
	version2 := "1.001"
	fmt.Println(compareVersion(version1, version2))
}
