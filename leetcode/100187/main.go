package main

import "fmt"

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	for i, j := c, d; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
		if i == a && j == b {
			break
		}
		if i == e && j == f {
			return 1
		}
		i++
		j++
	}

	for i, j := c, d; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
		if i == a && j == b {
			break
		}
		if i == e && j == f {
			return 1
		}
		i--
		j--
	}

	for i, j := c, d; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
		if i == a && j == b {
			break
		}
		if i == e && j == f {
			return 1
		}
		i++
		j--
	}

	for i, j := c, d; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
		if i == a && j == b {
			break
		}
		if i == e && j == f {
			return 1
		}
		i--
		j++
	}

	if (a == e) || (b == f) {
		for i, j := a, b; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
			if i == c && j == d {
				break
			}

			if i == e && j == f {
				return 1
			}

			i++
		}

		for i, j := a, b; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
			if i == c && j == d {
				break
			}

			if i == e && j == f {
				return 1
			}

			i--
		}

		for i, j := a, b; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
			if i == c && j == d {
				break
			}

			if i == e && j == f {
				return 1
			}

			j++
		}

		for i, j := a, b; i >= 0 && i <= 8 && j >= 0 && j <= 8; {
			if i == c && j == d {
				break
			}

			if i == e && j == f {
				return 1
			}

			j--
		}

	}

	return 2
}

func main() {
	a, b, c, d, e, f := 1, 2, 1, 4, 1, 6
	fmt.Println(minMovesToCaptureTheQueen(a, b, c, d, e, f))
}
