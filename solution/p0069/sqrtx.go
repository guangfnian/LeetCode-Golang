package p0069

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	l, r := 0, x
	var m int
	for l < r {
		m = (l + r) >> 1
		if m*m <= x {
			l = m + 1
		} else {
			r = m
		}
	}
	return r - 1
}
