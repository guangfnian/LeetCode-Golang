package p0397

func integerReplacement(n int) int {
	cnt := 0
	for n > 1 {
		if n&1 != 0 {
			if (n&3) == 3 && n != 3 {
				n++
			} else {
				n--
			}
		} else {
			n >>= 1
		}
		cnt++
	}
	return cnt
}
