package main

func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])

	all := make([][]byte, n+2)
	haveCome := make([][]bool, n+2)

	for i := 0; i < n+2; i++ {
		all = make([][]byte, m+2)
		haveCome = make([][]bool, m+2)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			all[i][j] = (grid[i][j])
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {

		}
	}
}
