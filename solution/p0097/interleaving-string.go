package p0097

func dfs(s1, s2, s3 string, c1, c2 int, mp *[][]int) bool {
	l1, l2 := len(s1), len(s2)
	c3 := c1 + c2
	if (*mp)[c1][c2] != 0 {
		if (*mp)[c1][c2] == 1 {
			return true
		} else {
			return false
		}
	}
	if l1 == c1 {
		if s2[c2:] == s3[c3:] {
			(*mp)[c1][c2] = 1
			return true
		} else {
			(*mp)[c1][c2] = 2
			return false
		}
	}
	if l2 == c2 {
		if s1[c1:] == s3[c3:] {
			(*mp)[c1][c2] = 1
			return true
		} else {
			(*mp)[c1][c2] = 2
			return false
		}
	}
	if s1[c1] == s3[c3] {
		if dfs(s1, s2, s3, c1+1, c2, mp) {
			(*mp)[c1][c2] = 1
			return true
		}
	}
	if s2[c2] == s3[c3] {
		if dfs(s1, s2, s3, c1, c2+1, mp) {
			(*mp)[c1][c2] = 1
			return true
		}
	}
	(*mp)[c1][c2] = 2
	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	var k = make([][]int, l1+1)
	for i := range k {
		k[i] = make([]int, l2+1)
	}
	return dfs(s1, s2, s3, 0, 0, &k)
}
