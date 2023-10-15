package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param donations int整型一维数组
 * @return int整型
 */
func maximizeDonations(donations []int) int {
	// write code here
	n := len(donations)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return donations[0]
	}

	if n%2 == 1 {
		minIndex := 0

		for i := range donations {
			if donations[i] < donations[minIndex] {
				minIndex = i
			}
		}

		if minIndex == n-1 {
			donations = donations[:n-1]
		} else {
			donations = append(donations[:minIndex], donations[minIndex+1:]...)
		}
	}

	sumJi := 0
	sumOu := 0

	fmt.Println(donations)

	for i := 0; i+1 < n; i = i + 2 {
		sumJi = sumJi + donations[i]
		sumOu = sumOu + donations[i+1]
	}

	return max(sumJi, sumOu)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximizeDonations1(donations []int) int {
	// write code here
	n := len(donations)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return donations[0]
	}

	start := 0
	if donations[0] < donations[1] {
		start = 1
	}

	sum := 0
	for i := start; i+3 < n+start-1; {
		fmt.Println(i)
		sum = sum + donations[i]
		if donations[i+2] >= donations[i+3] {
			fmt.Println(donations[i+2])
			sum = sum + donations[i+2]
			i = i + 2
		} else {
			fmt.Println(donations[i+3])
			sum = sum + donations[i+3]
			i = i + 3
		}
	}

	return sum
}

func main() {
	donations := []int{10, 3, 2, 5, 7, 8, 11}
	// fmt.Println(maximizeDonations(donations))
	fmt.Println(maximizeDonations1(donations))
}
