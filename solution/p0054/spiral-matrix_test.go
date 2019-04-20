package p0054

import (
	"reflect"
	"testing"
)

func TestOK(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
	}
	answers := [][]int{{1, 2, 3, 6, 9, 8, 7, 4, 5}, {1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}}
	for i := range inputs {
		if ret := spiralOrder(inputs[i]); !reflect.DeepEqual(ret, answers[i]) {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}
