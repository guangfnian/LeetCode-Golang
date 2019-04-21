package p0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	ret := make([][]int, 0)
	sort.Ints(nums)
	var l, r, sum int
	var i, j int
	for i = 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j = i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r = j+1, length-1
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < length-1 && nums[l] == nums[l+1] {
						l++
					}
					for r > 0 && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ret
}
