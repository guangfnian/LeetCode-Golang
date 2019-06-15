package p0059

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	vis := make([][]bool, n)
	for i := range ret {
		ret[i] = make([]int, n)
		vis[i] = make([]bool, n)
	}
	x, y := 0, 0
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	dir := 0
	target := n * n
	var nx, ny int
	for i := 1; i <= target; i++ {
		ret[x][y] = i
		vis[x][y] = true
		nx = x + dx[dir]
		ny = y + dy[dir]
		if nx < 0 || nx >= n || ny < 0 || ny >= n || vis[nx][ny] {
			dir = (dir + 1) % 4
			nx = x + dx[dir]
			ny = y + dy[dir]
		}
		x, y = nx, ny
	}
	return ret
}
