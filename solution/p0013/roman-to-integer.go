package p0013

func romanToInt(s string) int {
	mp := make([]int, 128)
	mp['I'], mp['V'], mp['X'], mp['L'], mp['C'], mp['D'], mp['M'] = 1, 5, 10, 50, 100, 500, 1000
	cur := 0
	ret := 0
	length := len(s)
	for cur < length {
		if cur < length-1 && mp[s[cur]] < mp[s[cur+1]] {
			ret += mp[s[cur+1]] - mp[s[cur]]
			cur += 2
		} else {
			ret += mp[s[cur]]
			cur++
		}
	}
	return ret
}
