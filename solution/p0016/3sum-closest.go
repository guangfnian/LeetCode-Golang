package p0016

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	diff := nums[0] + nums[1] + nums[2] - target
	cur := 0
	var l, r, sum int
	for cur < length {
		if cur > 0 && cur < length && nums[cur] == nums[cur-1] {
			cur++
		}
		l, r = cur+1, length-1
		for l < r {
			sum = nums[cur] + nums[l] + nums[r] - target
			if abs(sum) < abs(diff) {
				diff = sum
			}
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				return target
			}
		}
		cur++
	}
	return target + diff
}
