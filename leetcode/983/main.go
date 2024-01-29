package main

import "fmt"

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 365+1)

	m := make(map[int]struct{})
	for _, v := range days {
		m[v] = struct{}{}
	}

	day, week, month := costs[0], costs[1], costs[2]

	for i := 1; i < 366; i++ {
		if _, find := m[i]; find {
			dp[i] = dp[i-1] + day

			if i-7 >= 0 {
				dp[i] = min(dp[i-7]+week, dp[i])
			} else {
				dp[i] = min(dp[i], week)
			}

			if i-30 >= 0 {
				dp[i] = min(dp[i-30]+month, dp[i])
			} else {
				dp[i] = min(dp[i], month)
			}
		} else {
			dp[i] = dp[i-1]
		}

	}

	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	days := []int{1, 4, 6, 7, 8, 20}
	conts := []int{2, 7, 15}
	fmt.Println(mincostTickets(days, conts))
}
