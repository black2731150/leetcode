package main

import "sort"

func merge(A []int, m int, B []int, n int) {
	for i := m; i < m+n; i++ {
		A[i] = B[i-m]
	}

	sort.Ints(A)
}
