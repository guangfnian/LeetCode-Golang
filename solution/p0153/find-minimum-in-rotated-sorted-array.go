package p0153

func findMin(nums []int) int {
	length := len(nums)
	if nums[0] <= nums[length-1] {
		return nums[0]
	}
	l := 0
	r := length - 1
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
