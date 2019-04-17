package p0033

func binarySearch(nums []int, target, l, r int) int {
	if target < nums[l] || target > nums[r] {
		return -1
	}
	var m int
	for l < r {
		m = (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if nums[0] <= nums[length-1] {
		return binarySearch(nums, target, 0, length-1)
	}
	l := 0
	r := length - 1
	var m int
	for l < r {
		m = (l + r) >> 1
		if nums[m] > nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == 0 {
		l++
	}
	if target > nums[length-1] {
		return binarySearch(nums, target, 0, l-1)
	}
	return binarySearch(nums, target, l, length-1)
}
