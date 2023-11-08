package main

func permuteUnique(nums []int) [][]int {
	answer := [][]int{}

	generateNums(nums, []int{}, &answer)

	theMap := make(map[int][]int)
	for _, one := range answer {
		x := 0
		for _, n := range one {
			x = x*10 + n
		}
		theMap[x] = one
	}

	realAnswer := [][]int{}
	for _, v := range theMap {
		realAnswer = append(realAnswer, v)
	}
	return realAnswer
}

func generateNums(nums []int, dangqian []int, answer *[][]int) {
	if len(nums) == 0 {
		*answer = append(*answer, dangqian)
		return
	}

	for i, num := range nums {

		newNums := append([]int{}, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)

		generateNums(newNums, append(dangqian, num), answer)
	}
}
