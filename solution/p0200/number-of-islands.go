package p0200

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	vis := make([][]bool, n)
	for i := range grid {
		vis[i] = make([]bool, m)
	}
	ret := 0
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		vis[x][y] = true
		for i := range dx {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == '1' && !vis[nx][ny] {
				dfs(nx, ny)
			}
		}
	}
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '1' && !vis[i][j] {
				ret++
				dfs(i, j)
			}
		}
	}
	return ret
}
