package p0052

func dfs(cur, tar, n, ld, rd uint, ans *int) {
	if cur == tar {
		*ans++
		return
	}
	for i := uint(0); i < n; i++ {
		var t uint = 1 << i
		if cur&t == 0 && ld&t == 0 && rd&t == 0 {
			dfs(cur|t, tar, n, (ld|t)<<1, (rd|t)>>1, ans)
		}
	}
}

func totalNQueens(n int) int {
	ans := 0
	tar := (1 << uint(n)) - 1
	dfs(0, uint(tar), uint(n), 0, 0, &ans)
	return ans
}
