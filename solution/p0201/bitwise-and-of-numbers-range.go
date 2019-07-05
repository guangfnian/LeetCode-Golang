package p0201

func rangeBitwiseAnd(m int, n int) int {
	if n == m {
		return n
	}
	if m == 0 {
		return 0
	}
	l := uint32(32)
	for ; l > 0; l-- {
		if (m>>l)&1 != 0 {
			break
		}
	}

	r := uint32(32)
	for ; r > 0; r-- {
		if (n>>r)&1 != 0 {
			break
		}
	}
	if l != r {
		return 0
	}
	for ; l > 0; l-- {
		if (m>>l)&1 != (n>>l)&1 {
			break
		}
	}
	l++
	return (n >> l) << l
}
