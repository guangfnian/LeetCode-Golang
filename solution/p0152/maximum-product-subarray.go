package p0152

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	ret := nums[0]
	mx := nums[0]
	mn := nums[0]
	var t1, t2 int
	for i := 1; i < length; i++ {
		t1 = mx * nums[i]
		t2 = mn * nums[i]
		mx = max(nums[i], max(t1, t2))
		mn = min(nums[i], min(t1, t2))
		ret = max(ret, mx)
	}
	return ret
}
