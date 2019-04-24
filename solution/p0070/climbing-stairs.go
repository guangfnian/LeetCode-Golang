package p0070

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	a := 2
	b := 3
	for i := 3; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
