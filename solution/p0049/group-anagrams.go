package p0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := make(map[int][]string)
	hash := make(map[string]int)
	cur := 1
	var ret [][]string
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		ss := string(b)
		if h, ok := hash[ss]; !ok {
			hash[ss] = cur
			mp[cur] = append(mp[cur], s)
			cur++
		} else {
			mp[h] = append(mp[h], s)
		}
	}
	for _, v := range mp {
		ret = append(ret, v)
	}
	return ret
}
