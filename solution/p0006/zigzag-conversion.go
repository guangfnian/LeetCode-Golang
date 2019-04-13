package p0006

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	L := numRows<<1 - 2
	l := len(s)
	ret := make([]byte, l)
	sl, sr := L, L
	c := 0
	var cur int
	var op int
	for i := 0; i < numRows; i++ {
		cur = i
		op = 1
		for cur < l {
			ret[c] = s[cur]
			c++
			if op == 1 {
				cur += sl
			} else {
				cur += sr
			}
			op ^= 1
		}
		sl -= 2
		sr = L - sl
		if sl == 0 {
			sl, sr = L, L
		}
	}
	return string(ret)
}
