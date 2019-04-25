package p0091

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	if len(s) < 2 {
		return len(s)
	}
	dp := make([]int, len(s))
	dp[0], dp[1] = 1, 1
	if s[0] == '1' || (s[0] == '2' && s[1] < '7') {
		dp[1] = 2
	}
	if s[1] == '0' {
		if s[0] > '2' {
			return 0
		}
		dp[1] = 1
	}
	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '0' || s[i-1] > '2' {
				return 0
			}
			dp[i] = dp[i-2]
			continue
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}
