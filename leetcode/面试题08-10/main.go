package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]

	m := len(image)
	n := len(image[0])
	get := make([][]bool, m)
	for i := range get {
		get[i] = make([]bool, n)
	}

	type point struct {
		x int
		y int
	}
	quque := []point{{x: sr, y: sc}}
	for len(quque) > 0 {
		node := quque[0]
		quque = quque[1:]
		image[node.x][node.y] = newColor
		get[node.x][node.y] = true

		if node.x-1 >= 0 && image[node.x-1][node.y] == oldColor && !get[node.x-1][node.y] {
			quque = append(quque, point{x: node.x - 1, y: node.y})
		}

		if node.x+1 < m && image[node.x+1][node.y] == oldColor && !get[node.x+1][node.y] {
			quque = append(quque, point{x: node.x + 1, y: node.y})
		}

		if node.y-1 >= 0 && image[node.x][node.y-1] == oldColor && !get[node.x][node.y-1] {
			quque = append(quque, point{x: node.x, y: node.y - 1})
		}

		if node.y+1 < n && image[node.x][node.y+1] == oldColor && !get[node.x][node.y+1] {
			quque = append(quque, point{x: node.x, y: node.y + 1})
		}
	}
	return image
}
