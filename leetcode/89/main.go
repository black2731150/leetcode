package main

import "fmt"

func grayCode(n int) []int {
	answer := make([]int, 1<<n)
	themap := make(map[int]bool, 1<<n)
	themap[0] = true

	for i := 1; i <= 1<<n; i++ {
		before := answer[i-1]
		for j := 0; j < n; j++ {
			if x := before ^ (1 << j); !themap[x] {
				themap[x] = true
				answer[i] = x
				break
			}
		}
	}
	return answer

}

func main() {
	n := 2
	fmt.Println(grayCode(n))
}
