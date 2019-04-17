package p0034

func lowerBound(nums []int, target int) int {
	length := len(nums)
	l, r := 0, length-1
	var m int
	for l <= r {
		m = (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func upperBound(nums []int, target int) int {
	length := len(nums)
	l, r := 0, length-1
	var m int
	for l <= r {
		m = (l + r) >> 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	if nums[r] != target {
		return -1
	}
	return r
}

func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return ret
	}
	lower := lowerBound(nums, target)
	if lower == -1 {
		return ret
	}
	return []int{lower, upperBound(nums, target)}
}
