package p0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; ; i++ {
		for j := range strs {
			if i == len(strs[j]) {
				return strs[j][:i]
			}
			if j > 0 && strs[j][i] != strs[j-1][i] {
				return strs[j][:i]
			}
		}
	}
}
