package p0056

import "sort"

type Interval struct {
	Start int
	End   int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func merge(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start != intervals[j].Start {
			return intervals[i].Start < intervals[j].Start
		}
		return intervals[i].End < intervals[j].End
	})
	ans := make([]Interval, 0)
	if len(intervals) == 0 {
		return ans
	}
	l := intervals[0].Start
	r := intervals[0].End
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= r {
			r = max(r, intervals[i].End)
		} else {
			ans = append(ans, Interval{l, r})
			l = intervals[i].Start
			r = intervals[i].End
		}
	}
	ans = append(ans, Interval{l, r})
	return ans
}
