package main

func findErrorNums(nums []int) []int {
	var themap map[int]int = make(map[int]int, len(nums)+1)

	for i := 1; i < len(nums)+1; i++ {
		themap[i] = 0
	}

	for i := 0; i < len(nums); i++ {
		themap[nums[i]]++
	}

	var answer []int = make([]int, 0)
	var queshi, chongfu = 0, 0
	for i := 1; i < len(themap)+1; i++ {
		if themap[i] == 2 {
			chongfu = i
		}
		if themap[i] == 0 {
			queshi = i
		}
	}
	answer = append(answer, chongfu, queshi)

	return answer
}
