package main

//伟大的分差数组

func maximumBeauty(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	// 确定差分数组的大小
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	diff := make([]int, maxNum+k+2)

	// 更新差分数组
	for _, num := range nums {
		diff[max(0, num-k)]++
		diff[min(len(diff)-1, num+k+1)]--
	}

	// 计算最终的计数
	maxCount, currCount := 0, 0
	for _, d := range diff {
		currCount += d
		if currCount > maxCount {
			maxCount = currCount
		}
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//时间复杂度过高
// func maximumBeauty(nums []int, k int) int {
// 	numsMap := make(map[int]int)

// 	for _, v := range nums {
// 		for j := v - k; j <= v+k; j++ {
// 			numsMap[j]++
// 		}
// 	}

// 	type numsTime struct {
// 		num  int
// 		time int
// 	}

// 	slice := []numsTime{}
// 	for k2, v := range numsMap {
// 		slice = append(slice, numsTime{num: k2, time: v})
// 	}

// 	sort.Slice(slice, func(i, j int) bool {
// 		return slice[i].time > slice[j].time
// 	})

// 	return slice[0].time
// }
