package main

import "fmt"

//官方解答
// func numTimesAllBlue(flips []int) int {
//     n, ans, right := len(flips), 0, 0
//     for i := 0; i < n; i++ {
//         if flips[i] > right {
//             right = flips[i]
//         }
//         if right == i + 1 {
//             ans++
//         }
//     }
//     return ans
// }

func numTimesAllBlue(flips []int) int {
	n := len(flips)
	s := make([]bool, n)

	answer := 0

	lastOne := -1
	for _, v := range flips {
		v = v - 1

		s[v] = !s[v]

		if s[v] {
			if v > lastOne {
				lastOne = v
			}
		}

		x := true
		for j := lastOne; j >= 0; j-- {
			x = x && s[j]
			if !x {
				break
			}
		}

		if x {
			answer++
		}

	}

	return answer
}

func main() {
	flips := []int{3, 2, 4, 1, 5}
	fmt.Println(numTimesAllBlue(flips))
}
