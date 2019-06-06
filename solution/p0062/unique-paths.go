package p0062

func uniquePaths(m int, n int) int {
	m--
	n--
	ret := 1
	if n < m {
		n, m = m, n
	}
	for i := n + 1; i <= n+m; i++ {
		ret *= i
	}
	for i := 2; i <= m; i++ {
		ret /= i
	}
	return ret
}
