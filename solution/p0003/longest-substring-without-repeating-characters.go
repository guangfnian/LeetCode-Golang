package p0003

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	ans, l, r := 0, 0, 0
	count := make([]int, 128)
	for r < length {
		count[s[r]]++
		for count[s[r]] > 1 {
			count[s[l]]--
			l++
		}
		r++
		ans = max(ans, r-l)
	}
	return ans
}
