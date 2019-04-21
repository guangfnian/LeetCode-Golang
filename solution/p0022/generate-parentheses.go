package p0022

var ret []string

func dfs(x, y int, s []byte, cur int) {
	if x > y || x < 0 {
		return
	}
	if x == 0 && y == 0 {
		ret = append(ret, string(s))
		return
	}
	if x > 0 {
		s[cur] = '('
		dfs(x-1, y, s, cur+1)
	}
	if y > 0 {
		s[cur] = ')'
		dfs(x, y-1, s, cur+1)
	}
}

func generateParenthesis(n int) []string {
	ret = make([]string, 0)
	s := make([]byte, n<<1)
	dfs(n, n, s, 0)
	return ret
}
