package p0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := [][]int{{2, 7, 11, 15}}
	targets := []int{9}
	answers := [][]int{{0, 1}}
	length := len(nums)
	for i := 0; i < length; i++ {
		res := twoSum(nums[i], targets[i])
		if !reflect.DeepEqual(res, answers[i]) {
			t.Fatal(nums[i], targets[i], answers[i])
		}
	}
}