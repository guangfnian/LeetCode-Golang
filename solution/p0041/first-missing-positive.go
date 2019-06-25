package p0041

func firstMissingPositive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 1
	}
	for i := range nums {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return length + 1
}
