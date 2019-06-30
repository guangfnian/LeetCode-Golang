package p0079

func exist(board [][]byte, word string) bool {
	n := len(board)
	if n == 0 {
		return false
	}
	m := len(board[0])
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	length := len(word)
	if length == 0 {
		return false
	}
	flag := false
	var dfs func(x, y, cur int)
	dfs = func(x, y, cur int) {
		if flag {
			return
		}
		if cur == length {
			flag = true
			return
		}
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && board[nx][ny] == word[cur] && !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny, cur+1)
				vis[nx][ny] = false
			}
			if flag {
				return
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				vis[i][j] = true
				dfs(i, j, 1)
				if flag {
					break
				}
				vis[i][j] = false
			}
		}
	}
	return flag
}
