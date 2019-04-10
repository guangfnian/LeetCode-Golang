package p0011

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

func maxArea(height []int) int {
	length := len(height)
	l, r := 0, length-1
	ans := min(height[l], height[r]) * (r - l)
	for l < r {
		if height[l] < height[r] {
			l++
		} else if height[r] < height[l] {
			r--
		} else {
			l++
			r--
		}
		ans = max(ans, min(height[l], height[r])*(r-l))
	}
	return ans
}
