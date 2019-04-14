package p0015

import (
	"reflect"
	"testing"
)

func TestOK(t *testing.T) {
	inputs := [][]int{{-1, 0, 1, 2, -1, -4}, {0, 0}, {-2, 0, 1, 1, 2}, {0, 0, 0}}
	answers := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}, {}, {{-2, 0, 2}, {-2, 1, 1}}, {{0, 0, 0}}}
	for i := range inputs {
		if ret := threeSum(inputs[i]); !reflect.DeepEqual(ret, answers[i]) {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}
