package main

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		point := int64(questions[i][0])
		brainPower := questions[i][1]
		next := i + brainPower + 1
		if next < n {
			dp[i] = max(point+dp[next], dp[i+1])
		} else {
			dp[i] = max(point, dp[i+1])
		}
	}
	return dp[0]
}

// 递归方法 时间复杂度过高
// func mostPoints(questions [][]int) int64 {
// 	var help func(start int) int64
// 	help = func(start int) int64 {
// 		if start >= len(questions) {
// 			return 0
// 		} else {
// 			point, brainPower := questions[start][0], questions[start][1]
// 			return max(int64(point)+help(start+brainPower+1), help(start+1))
// 		}
// 	}
// 	return help(0)
// }

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	questions := [][]int{{3, 2}, {3, 4}, {4, 5}, {2, 5}}
}
