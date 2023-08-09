package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	//二分查找法
	l, r := 0, n
	for l <= r {
		mid := (l + r) / 2
		num := guess(mid)
		if num == 0 {
			return mid
		} else if num == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func guess(num int) int {
	Gussnum := 6
	if num == Gussnum {
		return 0
	}

	if num > Gussnum {
		return 1
	} else {
		return -1
	}

}

func main() {
	num := 10
	fmt.Println(guessNumber(num))
}
