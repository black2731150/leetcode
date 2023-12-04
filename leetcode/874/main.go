package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	fangs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	type point struct {
		x int
		y int
	}
	theMap := make(map[point]struct{})
	for _, v := range obstacles {
		theMap[point{x: v[0], y: v[1]}] = struct{}{}
	}

	fangxiang := fangs[0]
	fangsIndex := 0
	x, y := 0, 0
	answer := 0
	for _, v := range commands {
		if v > 0 {
			if fangxiang[0] == 0 {
				for j := 0; j < v; j++ {
					if _, find := theMap[point{x: x, y: y + fangxiang[1]*1}]; find {
						break
					} else {
						y = y + fangxiang[1]*1
						answer = max(answer, x*x+y*y)
					}
				}
				continue
			}

			if fangxiang[1] == 0 {
				for j := 0; j < v; j++ {
					if _, find := theMap[point{x: x + fangxiang[0]*1, y: y}]; find {
						break
					} else {
						x = x + fangxiang[0]*1
						answer = max(answer, x*x+y*y)
					}
				}
				continue
			}
		} else {
			if v == -2 {
				fangsIndex = (fangsIndex - 1) % 4
				if fangsIndex < 0 {
					fangsIndex += 4
				}
				fangxiang = fangs[fangsIndex]
			}

			if v == -1 {
				fangsIndex = (fangsIndex + 1) % 4
				fangxiang = fangs[fangsIndex]
			}
		}
	}

	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	commands := []int{6, -1, -1, 6}
	obstacles := [][]int{}
	fmt.Println(robotSim(commands, obstacles))
}
