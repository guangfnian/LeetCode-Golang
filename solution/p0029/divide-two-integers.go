package p0029

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func divide(dividend int, divisor int) int {
	const MAX = 1<<31 - 1
	const MIN = -1 << 31
	if dividend == MIN && divisor == -1 {
		return MAX
	}
	if dividend == 0 {
		return 0
	}
	ans := 0
	flag := (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	cache := make([]int, 0)
	twos := make([]int, 0)
	two := 1
	for divisor <= dividend {
		cache = append(cache, divisor)
		twos = append(twos, two)
		two = two + two
		divisor = divisor + divisor
	}
	cur := len(cache) - 1
	for cur >= 0 {
		for dividend >= cache[cur] {
			dividend -= cache[cur]
			ans += twos[cur]
		}
		cur--
	}
	if flag {
		ans = -ans
	}
	return ans
}
