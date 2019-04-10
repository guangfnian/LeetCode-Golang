package p0005

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestPalindrome(s string) string {
	b := make([]byte, (len(s)<<1)+1)
	b[0] = '#'
	for i := range s {
		b[(i<<1)+1] = byte(s[i])
		b[(i<<1)+2] = '#'
	}
	l, r := 0, 0
	pos, mr, maxLength := 0, 0, 0
	length := len(b)
	rl := make([]int, length)
	for i := range b {
		if i < mr {
			rl[i] = min(mr-i, rl[(pos<<1)-i])
		} else {
			rl[i] = 1
		}
		for i-rl[i] >= 0 && i+rl[i] < length && b[i-rl[i]] == b[i+rl[i]] {
			rl[i]++
		}
		if i+rl[i]-1 > mr {
			mr = i + rl[i] - 1
			pos = i
		}
		if rl[i] > maxLength {
			maxLength = rl[i]
			l = i - rl[i] + 1
			r = i + rl[i] - 1
		}
	}
	l = l >> 1
	r = r >> 1
	return s[l:r]
}
