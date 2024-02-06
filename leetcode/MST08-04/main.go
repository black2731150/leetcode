package main

func subsets(nums []int) [][]int {
	n := len(nums)

	answer := [][]int{}

	tmp := []int{}
	var help func(start int, k int)
	help = func(start int, k int) {

		if len(tmp) == k {
			ct := make([]int, len(tmp))
			copy(ct, tmp)
			answer = append(answer, ct)
			return
		}

		if start < 0 || start >= n {
			return
		}

		for i := start; i < n; i++ {
			if i > 0 && nums[i-1] == nums[i] {
				continue
			}
			tmp = append(tmp, nums[i])
			help(i+1, k)
			tmp = tmp[:len(tmp)-1]
		}
	}

	for i := 0; i <= n; i++ {
		tmp = []int{}
		help(0, i)
	}

	return answer
}
