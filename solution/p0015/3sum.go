package p0015

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	length := len(nums)
	var l, r, sum int
	cur := 0
	for cur < length {
		if nums[cur] > 0 {
			break
		}
		for cur > 0 && cur < length && nums[cur] == nums[cur-1] {
			cur++
		}
		l, r = cur+1, length-1
		for l < r {
			sum = nums[cur] + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ret = append(ret, []int{nums[cur], nums[l], nums[r]})
				for l < length-1 && nums[l] == nums[l+1] {
					l++
				}
				for r > 0 && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
		cur++
	}
	return ret
}

func threeSum2(nums []int) [][]int {
	mp := make(map[int]int)
	ret := make([][]int, 0)
	for i := range nums {
		mp[nums[i]]++
	}

	ns := make([]int, len(mp))
	cur := 0
	for i := range mp {
		ns[cur] = i
		cur++
	}
	sort.Ints(ns)
	l := len(ns)

	for i := range ns {
		v1, _ := mp[ns[i]]
		if ns[i] == 0 {
			if v1 > 2 {
				ret = append(ret, []int{0, 0, 0})
			}
			continue
		}
		if v1 > 1 {
			if _, ok := mp[-ns[i]-ns[i]]; ok {
				ret = append(ret, []int{ns[i], ns[i], -ns[i] - ns[i]})
			}
		}
		for j := i + 1; j < l; j++ {
			k3 := -ns[j] - ns[i]
			if k3 <= ns[j] {
				continue
			}
			if _, ok := mp[k3]; ok {
				ret = append(ret, []int{ns[i], ns[j], k3})
			}
		}
	}
	return ret
}
