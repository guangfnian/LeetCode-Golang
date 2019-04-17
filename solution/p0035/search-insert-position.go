package p0035

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
