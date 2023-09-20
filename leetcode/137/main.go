package main

func singleNumber(nums []int) int {
	var answer uint32 = 0

	for i := 0; i < 32; i++ {
		var all int32 = 0

		for j := range nums {
			all = all + ((int32(nums[j]) >> i) & 1)
		}

		if all%3 > 0 {
			answer = answer | (1 << i)
		}
	}

	return int(answer)
}
