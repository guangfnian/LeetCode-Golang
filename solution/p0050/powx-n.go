package p0050

func pow(x float64, n int) float64 {
	base := x
	ret := 1.0
	for n > 0 {
		if (n & 1) == 1 {
			ret *= base
		}
		base *= base
		n >>= 1
	}
	return ret
}

func myPow(x float64, n int) float64 {
	flag := false
	if n < 0 {
		flag = true
		n = -n
	}
	res := pow(x, n)
	if flag {
		res = 1.0 / res
	}
	return res
}
