package p0038

import "strconv"

func countAndSay(n int) string {
	ret := "1"
	var tmp string
	var l, r, length int
	for i := 2; i <= n; i++ {
		tmp = ""
		length = len(ret)
		l, r = 0, 0
		for r < length {
			for r < length && ret[r] == ret[l] {
				r++
			}
			tmp += strconv.Itoa(r-l) + string(ret[l])
			l = r
		}
		ret = tmp
	}
	return ret
}
