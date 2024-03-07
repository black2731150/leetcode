package main

/**
 * Note: 类名、方法名、参数名已经指定，请勿修改
 *
 *
 * 输入一个电话号码字符串，和一个字符串数组。判断字符串数组中，每个字符串能否通过【连续字符从小到大排序】的操作，整个字符串变得与给定的电话号码字符串相同，并将判定结果依次写入结果数组
 * @param phoneNum string字符串  电话号码，纯数字
 * @param cardListArray string字符串 一维数组 字符串数组，卡片序列组成的数组
 * @return int整型一维数组
 */
func SortSubStringToBuildPhoneNum(phoneNum string, cardListArray []string) []int {
	phoneMap := make(map[byte]int)
	for i := range phoneNum {
		phoneMap[phoneNum[i]]++
	}

	answer := make([]int, len(cardListArray))
	for cardIndex, card := range cardListArray {
		cardMap := make(map[byte]int)
		haveLetter := false
		for i := range card {
			cardMap[card[i]]++
			if (card[i] >= 'a' && card[i] <= 'z') || (card[i] >= 'A' && card[i] <= 'Z') {
				haveLetter = true
			}

			if haveLetter {
				answer[cardIndex] = -1
			} else {
				ok := true
				for b1, num1 := range phoneMap {
					if num2, find := cardMap[b1]; !find || num1 != num2 {
						answer[cardIndex] = 0
						ok = false
						break
					}
				}

				if ok {
					answer[cardIndex] = 1
				}
			}
		}
	}

	return answer
}
