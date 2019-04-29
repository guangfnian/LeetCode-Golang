package p0130

func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])
	vis := make([][]bool, n)
	for i := range board {
		vis[i] = make([]bool, m)
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	var f func(x, y int)
	f = func(x, y int) {
		vis[x][y] = true
		for i := range dx {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && board[nx][ny] == 'O' && !vis[nx][ny] {
				f(nx, ny)
			}
		}
	}
	for i := range board {
		if board[i][0] == 'O' {
			f(i, 0)
		}
		if board[i][m-1] == 'O' {
			f(i, m-1)
		}
	}
	for i := range board[0] {
		if board[0][i] == 'O' {
			f(0, i)
		}
		if board[n-1][i] == 'O' {
			f(n-1, i)
		}
	}
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' && !vis[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
