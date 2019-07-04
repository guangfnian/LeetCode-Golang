package p0162

func findPeakElement(nums []int) int {
	length := len(nums)
	if length <= 1 || nums[0] > nums[1] {
		return 0
	}
	if nums[length-1] > nums[length-2] {
		return length - 1
	}
	l := 0
	r := length - 1
	var m int
	for l < r {
		m = (l + r) >> 1
		if m > 0 && m < length && nums[m] > nums[m-1] && nums[m] > nums[m+1] {
			return m
		}
		if m > 0 && nums[m] < nums[m-1] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
