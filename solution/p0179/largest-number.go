package p0179

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i := range nums {
		ss[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] > ss[j]+ss[i]
	})
	ret := ""
	for i := range ss {
		ret += ss[i]
	}
	for ret[0] == '0' && len(ret) > 1 {
		ret = ret[1:]
	}
	return ret
}
