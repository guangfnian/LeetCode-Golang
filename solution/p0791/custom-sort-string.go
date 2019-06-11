package p0791

import "sort"

func customSortString(S string, T string) string {
	mp := make([]int, 128)
	cur := 1
	for i := range S {
		mp[S[i]] = cur
		cur++
	}
	b := []byte(T)
	sort.Slice(b, func(i, j int) bool {
		return mp[b[i]] < mp[b[j]]
	})
	return string(b)
}
