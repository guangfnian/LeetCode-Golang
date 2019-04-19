package p0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := make([]int, 11)
	cur := 0
	for x > 9 {
		arr[cur] = x % 10
		cur++
		x /= 10
	}
	arr[cur] = x
	l, r := 0, cur
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
