package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	theMap := make(map[int]int)
	for _, v := range nums {
		theMap[v]++
	}

	type Point struct {
		num   int
		times int
	}

	Points := []Point{}
	for k, v := range theMap {
		Points = append(Points, Point{num: k, times: v})
	}

	sort.Slice(Points, func(i, j int) bool {
		return Points[i].num < Points[j].num
	})

	if len(Points) == 1 {
		return Points[0].num * Points[0].times
	}

	bp := make([]int, len(Points))
	bp[0] = Points[0].num * Points[0].times
	if Points[1].num-1 != Points[0].num {
		bp[1] = bp[0] + Points[1].num*Points[1].times
	} else {
		bp[1] = max(bp[0], Points[1].num*Points[1].times)
	}

	for i := 2; i < len(Points); i++ {
		if Points[i].num-1 == Points[i-1].num {
			bp[i] = max(bp[i-2]+Points[i].num*Points[i].times, bp[i-1])
		} else {
			bp[i] = bp[i-1] + Points[i].num*Points[i].times
		}
	}

	return bp[len(bp)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{2, 2, 3, 3, 3, 4}
	fmt.Println(deleteAndEarn(nums))
}
