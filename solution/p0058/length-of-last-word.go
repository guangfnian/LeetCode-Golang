package p0058

func lengthOfLastWord(s string) int {
	length := len(s)
	r := length - 1
	for r >= 0 && s[r] == ' ' {
		r--
	}
	l := r
	for l >= 0 && s[l] != ' ' {
		l--
	}
	return r - l
}
