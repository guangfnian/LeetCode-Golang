package p0051

var ret [][]string

func dfs(cur, tar, n, ld, rd, row uint, tmp, str []string) {
	if cur == tar {
		cp := make([]string, n)
		copy(cp, tmp)
		ret = append(ret, cp)
		return
	}
	for i := uint(0); i < n; i++ {
		var t uint = 1 << i
		if cur&t == 0 && ld&t == 0 && rd&t == 0 {
			tmp[row] = str[i]
			dfs(cur|t, tar, n, (ld|t)<<1, (rd|t)>>1, row+1, tmp, str)
		}
	}
}

func solveNQueens(n int) [][]string {
	strs := make([]string, n)
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = '.'
	}
	for i := 0; i < n; i++ {
		bytes[i] = 'Q'
		strs[i] = string(bytes)
		bytes[i] = '.'
	}
	ret = make([][]string, 0)
	tar := (1 << uint(n)) - 1
	tmp := make([]string, n)
	dfs(0, uint(tar), uint(n), 0, 0, 0, tmp, strs)
	return ret
}
