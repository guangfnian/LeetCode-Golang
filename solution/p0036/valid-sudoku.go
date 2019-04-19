package p0036

const N = 9

func checkRow(board [][]byte, n int) bool {
	cnt := make([]bool, 128)
	for i := 0; i < N; i++ {
		if board[n][i] == '.' {
			continue
		}
		if cnt[board[n][i]] {
			return false
		}
		cnt[board[n][i]] = true
	}
	return true
}

func checkColumn(board [][]byte, n int) bool {
	cnt := make([]bool, 128)
	for i := 0; i < N; i++ {
		if board[i][n] == '.' {
			continue
		}
		if cnt[board[i][n]] {
			return false
		}
		cnt[board[i][n]] = true
	}
	return true
}

func checkGrid(board [][]byte, x, y int) bool {
	cnt := make([]bool, 128)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if cnt[board[i][j]] {
				return false
			}
			cnt[board[i][j]] = true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < N; i++ {
		if !checkRow(board, i) || !checkColumn(board, i) {
			return false
		}
	}
	for i := 0; i < N; i += 3 {
		for j := 0; j < N; j += 3 {
			if !checkGrid(board, i, j) {
				return false
			}
		}
	}
	return true
}
