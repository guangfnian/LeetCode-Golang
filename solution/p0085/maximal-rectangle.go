package p0085

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	rm := make([]int, m)
	lm := make([]int, m)

	var maxArea = func(h []int) int {
		rm[m-1] = 1
		var j int
		for i := m - 2; i >= 0; i-- {
			if h[i] > h[i+1] {
				rm[i] = 1
			} else {
				j = i + 1
				for j < m && h[j] >= h[i] {
					j += rm[j]
				}
				rm[i] = j - i
			}
		}

		lm[0] = 1
		for i := 1; i < m; i++ {
			if h[i] > h[i-1] {
				lm[i] = 1
			} else {
				j = i - 1
				for j >= 0 && h[j] >= h[i] {
					j -= lm[j]
				}
				lm[i] = i - j
			}
		}
		r := 0
		for i := 0; i < m; i++ {
			r = max(r, h[i]*(lm[i]+rm[i]-1))
		}
		return r
	}

	count := make([][]int, n)
	for i := range matrix {
		count[i] = make([]int, m)
	}
	for i := range matrix[0] {
		if matrix[0][i] == '1' {
			count[0][i] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := range matrix[0] {
			if matrix[i][j] == '1' {
				count[i][j] = count[i-1][j] + 1
			}
		}
	}
	ret := 0
	for i := range count {
		ret = max(ret, maxArea(count[i]))
	}
	return ret
}
