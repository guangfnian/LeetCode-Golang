package p0004

import "math"

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

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n, m := len(nums1), len(nums2)
	var lm1, lm2, rm1, rm2, c1, c2 int
	tot := n + m
	n <<= 1
	m <<= 1
	l, r := 0, n
	for l <= r {
		c1 = (l + r) >> 1
		c2 = tot - c1

		if c1 == 0 {
			lm1 = math.MinInt32
		} else {
			lm1 = nums1[(c1-1)>>1]
		}

		if c1 == n {
			rm1 = math.MaxInt32
		} else {
			rm1 = nums1[c1>>1]
		}

		if c2 == 0 {
			lm2 = math.MinInt32
		} else {
			lm2 = nums2[(c2-1)>>1]
		}

		if c2 == m {
			rm2 = math.MaxInt32
		} else {
			rm2 = nums2[c2>>1]
		}

		if lm1 > rm2 {
			r = c1 - 1
		} else if lm2 > rm1 {
			l = c1 + 1
		} else {
			break
		}
	}
	return float64(max(lm1, lm2)+min(rm1, rm2)) / 2.0
}
