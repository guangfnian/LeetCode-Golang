package p0046

import "sort"

func reverse(nums []int, x, y int) {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) bool {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := n - 1; ; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					reverse(nums, i, n-1)
					return true
				}
			}
		}
	}
	return false
}
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for {
		tmp := make([]int, n)
		copy(tmp, nums)
		ret = append(ret, tmp)
		if !nextPermutation(nums) {
			break
		}
	}
	return ret
}
