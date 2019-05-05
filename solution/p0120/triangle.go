package p0120

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][i] += triangle[i-1][i-1]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
		}
	}
	ret := triangle[n-1][0]
	for i := range triangle[n-1] {
		ret = min(ret, triangle[n-1][i])
	}
	return ret
}
