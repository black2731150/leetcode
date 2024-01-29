package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for _, v := range image {
		for i2, j := 0, len(v)-1; i2 < j; i2, j = i2+1, j-1 {
			v[i2], v[j] = v[j], v[i2]
		}

		for i := range v {
			if v[i] == 1 {
				v[i] = 0
			} else {
				v[i] = 1
			}
		}
	}
	return image
}

func main() {
	image := [][]int{{1}}
	fmt.Println(flipAndInvertImage(image))
}
