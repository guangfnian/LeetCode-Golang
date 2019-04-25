package p0078

import "math/bits"

func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	length := uint(len(nums))
	total := uint(1 << length)
	var cur int
	for i := uint(0); i < total; i++ {
		ans := make([]int, bits.OnesCount(i))
		cur = 0
		for j := uint(0); j < length; j++ {
			if i&(1<<j) != 0 {
				ans[cur] = nums[j]
				cur++
			}
		}
		ret = append(ret, ans)
	}
	return ret
}
