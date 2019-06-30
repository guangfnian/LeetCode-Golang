package p0084

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}
	rm := make([]int, length)
	rm[length-1] = 1
	var j int
	for i := length - 2; i >= 0; i-- {
		if heights[i] > heights[i+1] {
			rm[i] = 1
		} else {
			j = i + 1
			for j < length && heights[j] >= heights[i] {
				j += rm[j]
			}
			rm[i] = j - i
		}
	}

	lm := make([]int, length)
	lm[0] = 1
	for i := 1; i < length; i++ {
		if heights[i] > heights[i-1] {
			lm[i] = 1
		} else {
			j = i - 1
			for j >= 0 && heights[j] >= heights[i] {
				j -= lm[j]
			}
			lm[i] = i - j
		}
	}
	ans := 0
	for i := 0; i < length; i++ {
		ans = max(ans, heights[i]*(rm[i]+lm[i]-1))
	}
	return ans
}
