package p0037

var dt = []int{0, 0, 0, 3, 3, 3, 6, 6, 6}

const N = 9

var flag bool

func isValid(board [][]byte, x, y int, p byte) bool {
	for i := 0; i < N; i++ {
		if board[x][i] == p || board[i][y] == p {
			return false
		}
		if board[dt[x]+i/3][dt[y]+i%3] == p {
			return false
		}
	}
	return true
}

func dfs(board [][]byte, pos int) {
	if flag {
		return
	}
	if pos > 80 {
		flag = true
		return
	}
	x := pos / 9
	y := pos % 9
	if board[x][y] != '.' {
		dfs(board, pos+1)
		return
	}
	for i := byte('1'); i <= '9'; i++ {
		if isValid(board, x, y, i) {
			board[x][y] = i
			dfs(board, pos+1)
		}
		if flag {
			return
		}
	}
	board[x][y] = '.'
}

func solveSudoku(board [][]byte) {
	flag = false
	dfs(board, 0)
}
