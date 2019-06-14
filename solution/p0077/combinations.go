package p0077

import "math/bits"

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	for i := 0; i < (1 << uint(n)); i++ {
		if bits.OnesCount32(uint32(i)) == k {
			var cur []int
			for j := 0; j < n; j++ {
				if i&(1<<uint(j)) != 0 {
					cur = append(cur, j+1)
				}
			}
			ret = append(ret, cur)
		}
	}
	return ret
}
