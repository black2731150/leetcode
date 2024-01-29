package main

type myHeap []int

func (m myHeap) Len() int           { return len(m) }
func (m myHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m myHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *myHeap) Pop() any {
	answer := (*m)[0]
	(*m) = (*m)[1:]
	return answer
}

func (m *myHeap) Push(v any) {
	*m = append(*m, v.(int))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	answer := [][]int{}
	for len(nums1) > 0 && k > 0 {
		top := nums1[0]
		second := 0
		nums1 = nums1[1:]

		if len(nums1) > 0 {
			second = nums1[0]
		}

		for _, v := range nums2 {
			if top+v < second+v {
				answer = append(answer, []int{top, v})
				k--
			} else {
				break
			}

		}
	}

	return answer
}
