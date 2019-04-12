package p0007

func reverse(x int) int {
	const Max = 1<<31 - 1
	const Min = -1 << 31
	flag := x < 0
	if flag {
		x = -x
	}
	xx := int64(x)
	var ret int64 = 0
	for xx > 9 {
		ret = ret*10 + xx%10
		xx /= 10
	}
	ret = ret*10 + xx
	if flag {
		ret = -ret
	}
	if ret > Max || ret < Min {
		return 0
	}
	return int(ret)
}
