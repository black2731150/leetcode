package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, v := range flowerbed {
		if v == 0 {
			if i > 0 && i < len(flowerbed)-1 {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else if i == 0 {
				if len(flowerbed) > 1 {
					if flowerbed[1] == 0 {
						flowerbed[i] = 1
						n--
					}
				} else {
					flowerbed[i] = 1
					n--
				}
			} else if i == len(flowerbed)-1 {
				if len(flowerbed) > 1 {
					if flowerbed[len(flowerbed)-2] == 0 {
						flowerbed[i] = 1
						n--
					}
				} else {
					flowerbed[i] = 1
					n--
				}
			}
		}
	}

	return n <= 0
}

func main() {
	flowerbed := []int{0, 0, 1, 0, 0}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
