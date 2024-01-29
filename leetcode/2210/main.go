package main

import "fmt"

func countHillValley(nums []int) int {
	answer := 0
	help := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if help[len(help)-1] == nums[i] {
			continue
		} else {
			help = append(help, nums[i])
		}
	}

	for i := 1; i < len(help)-1; {
		if help[i] > help[i-1] && help[i] > help[i+1] {
			answer++
			i++
			continue
		}

		if help[i] < help[i-1] && help[i] < help[i+1] {
			answer++
			i++
			continue
		}
		i++
	}
	return answer
}

func main() {
	nums := []int{2, 4, 1, 1, 6, 5}
	fmt.Println(countHillValley(nums))
}
