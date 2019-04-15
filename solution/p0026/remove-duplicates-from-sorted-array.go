package p0026

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	cur := 0
	r := 0
	for r < length {
		for r < length && nums[r] == nums[cur] {
			r++
		}
		cur++
		if r < length {
			nums[cur] = nums[r]
		}
	}
	return cur
}
