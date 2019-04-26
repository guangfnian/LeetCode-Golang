package p0080

func removeDuplicates(nums []int) int {
	cur := 0
	pos := 0
	length := len(nums)
	for pos < length {
		if cur > 1 && nums[cur-1] == nums[cur-2] {
			for pos < length && nums[pos] == nums[cur-1] {
				pos++
			}
		}
		if pos < length {
			nums[cur] = nums[pos]
			cur++
		}
		pos++
	}
	return cur
}
