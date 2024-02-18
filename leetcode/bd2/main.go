package main

import "fmt"

func main() {
	N, D := 0, 0
	fmt.Scanln(&N, &D)

	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(f(N, D, nums))
}

func f(N, D int, nums []int) int {
	if N <= 2 {
		return 0
	}

	mod := 99997867

	answer := 0

	for i := 0; i < N; i++ {
		start := nums[i]
		end := start + D

		left, right := i, N-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == end {
				left = mid
				break
			}

			if nums[mid] > end {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		tmp := 0
		if left == N {
			left = N - 1
			tmp = N - i
		} else {
			tmp = left - i
		}

		if nums[left] != end {
			tmp--
		}

		// fmt.Printf("tmp: %d\n", tmp)

		if tmp > 0 {
			answer = (answer + CN2(tmp)) % mod
		}

	}

	return answer
}

func CN2(n int) int {
	return n * (n - 1) / 2
}
