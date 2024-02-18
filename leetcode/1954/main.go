package main

var nums []int64 = []int64{0}

func init() {
	for i := 1; nums[len(nums)-1] < int64(1e15); i++ {
		nums = append(nums, 0)
		nums[i] = nums[i-1] + int64(4*i*(i+1))
	}
}

func minimumPerimeter(neededApples int64) int64 {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == neededApples {
			return int64(mid) * 8
		}

		if nums[mid] < neededApples {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return int64(left) * 8
}

func main() {

}
