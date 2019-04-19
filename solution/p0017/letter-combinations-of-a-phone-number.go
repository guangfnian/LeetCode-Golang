package p0017

var ret []string

var c = [][]byte{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func dfs(digits string, cur int, now []byte) {
	if cur == len(digits) {
		ret = append(ret, string(now))
		return
	}
	d := digits[cur] - '2'
	for i := range c[d] {
		now[cur] = c[d][i]
		dfs(digits, cur+1, now)
	}
}

func letterCombinations(digits string) []string {
	ret = make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	now := make([]byte, len(digits))
	dfs(digits, 0, now)
	return ret
}
