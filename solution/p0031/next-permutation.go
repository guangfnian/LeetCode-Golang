package p0031

func reverse(nums []int, x, y int) {
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := n - 1; ; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					reverse(nums, i, n-1)
					return
				}
			}
		}
	}
	reverse(nums, 0, n-1)
}
