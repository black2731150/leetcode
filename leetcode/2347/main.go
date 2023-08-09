package main

import (
	"fmt"
	"sort"
)

func bestHand(ranks []int, suits []byte) string {

	sort.Ints(ranks)

	if isFlush(suits) {
		return "Flush"
	}

	if isThreeOfAKind(ranks) {
		return "Three of a Kind"
	}

	if isPair(ranks) {
		return "Pair"
	}

	if isHighCard(ranks) {
		return "High Card"
	}

	return ""

}

func isFlush(suits []byte) bool {
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			return false
		}
	}
	return true
}

func isThreeOfAKind(nums []int) bool {
	for i := 2; i < len(nums); i++ {
		if (nums[i-2] == nums[i-1]) && (nums[i-1] == nums[i]) {
			return true
		}
	}
	return false
}

func isPair(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}

func isHighCard(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return false
		}
	}
	return true
}

func main() {
	ranks := []int{1, 2, 3, 4, 5}
	suits := []byte{'a', 'a', 'a', 'a', 'a'}
	fmt.Println(bestHand(ranks, suits))
}
