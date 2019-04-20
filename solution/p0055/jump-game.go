package p0055

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	r := 0
	for i := 0; i < length; i++ {
		if i > r {
			return false
		}
		r = max(r, i+nums[i])
		if r >= length-1 {
			return true
		}
	}
	return false
}
