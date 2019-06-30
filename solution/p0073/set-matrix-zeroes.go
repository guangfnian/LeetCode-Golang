package p0073

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	fn, fm := false, false
	for i := range matrix {
		if matrix[i][0] == 0 {
			fm = true
			break
		}
	}
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			fn = true
			break
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < m; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < n; j++ {
				matrix[j][i] = 0
			}
		}
	}
	if fn {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	if fm {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
