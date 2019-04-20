package p0054

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	m := len(matrix[0])
	length := n * m
	ret := make([]int, length)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	x, y := 0, 0
	dir := 0
	for i := 0; i < length; i++ {
		ret[i] = matrix[x][y]
		vis[x][y] = true
		nx := x + dx[dir]
		ny := y + dy[dir]
		if nx < 0 || nx >= n || ny < 0 || ny >= m || vis[nx][ny] {
			dir = (dir + 1) % 4
			x += dx[dir]
			y += dy[dir]
		} else {
			x = nx
			y = ny
		}
	}
	return ret
}
