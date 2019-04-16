package p0027

func removeElement(nums []int, val int) int {
	length := len(nums)
	cur := 0
	r := 0
	for {
		for r < length && nums[r] == val {
			r++
		}
		if r == length {
			break
		}
		nums[cur] = nums[r]
		cur++
		r++
	}
	return cur
}
