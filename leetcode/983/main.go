package main

func mincostTickets(days []int, costs []int) int {
	chufa := make([]bool, 365)
	for _, v := range days {
		chufa[v-1] = true
	}

	one := costs[0]
	seven := costs[1]
	thrty := costs[2]

	bp := make([]int, len(days))
	bp[0] = one

	sevenDay := 0
	for i := 0; i < 7; i++ {
		if chufa[i] {
			sevenDay++
		}
	}
	if seven <= sevenDay*one {
		bp[0] = seven
		for i := 0; i < 7; i++ {
			chufa[i] = false
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
